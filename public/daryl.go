package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/config"
	"github.com/vitaminwater/daryl/protodef"
)

const AUTH_TOKEN_HEADER = "X-Daryl-Auth-Token"

type darylCommand interface {
	Name() string
	Object() interface{}
	Execute(*gin.Context, protodef.DarylServiceClient, interface{}) (interface{}, error)
}

type userMessageCommand struct {
}

func (c *userMessageCommand) Name() string {
	return "message"
}

func (c *userMessageCommand) Object() interface{} {
	return &protodef.Message{}
}

func (c *userMessageCommand) Execute(co *gin.Context, d protodef.DarylServiceClient, o interface{}) (interface{}, error) {
	i := co.MustGet("daryl_id").(string)
	um := protodef.UserMessageRequest{DarylIdentifier: i, Message: o.(*protodef.Message)}
	return d.UserMessage(context.Background(), &um)
}

type addHabitCommand struct {
}

func (c *addHabitCommand) Name() string {
	return "habit"
}

func (c *addHabitCommand) Object() interface{} {
	return &protodef.Habit{}
}

func (c *addHabitCommand) Execute(co *gin.Context, d protodef.DarylServiceClient, o interface{}) (interface{}, error) {
	i := co.MustGet("daryl_id").(string)
	log.Info(o)
	ah := protodef.AddHabitRequest{DarylIdentifier: i, Habit: (o.(*protodef.Habit))}
	return d.AddHabit(context.Background(), &ah)
}

type startWorkSessionCommand struct {
}

func (c *startWorkSessionCommand) Name() string {
	return "session"
}

func (c *startWorkSessionCommand) Object() interface{} {
	return &protodef.SessionConfig{}
}

func (c *startWorkSessionCommand) Execute(co *gin.Context, d protodef.DarylServiceClient, o interface{}) (interface{}, error) {
	i := co.MustGet("daryl_id").(string)
	r := protodef.StartWorkSessionRequest{DarylIdentifier: i, Config: o.(*protodef.SessionConfig)}
	return d.StartWorkSession(context.Background(), &r)
}

type cancelWorkSessionCommand struct {
}

func (c *cancelWorkSessionCommand) Name() string {
	return "cancel"
}

func (c *cancelWorkSessionCommand) Object() interface{} {
	return nil
}

func (c *cancelWorkSessionCommand) Execute(co *gin.Context, d protodef.DarylServiceClient, o interface{}) (interface{}, error) {
	i := co.MustGet("daryl_id").(string)
	r := protodef.CancelWorkSessionRequest{DarylIdentifier: i}
	return d.CancelWorkSession(context.Background(), &r)
}

type refuseSessionSliceCommand struct {
}

func (c *refuseSessionSliceCommand) Name() string {
	return "refuse"
}

func (c *refuseSessionSliceCommand) Object() interface{} {
	return &protodef.SessionSliceIndex{}
}

func (c *refuseSessionSliceCommand) Execute(co *gin.Context, d protodef.DarylServiceClient, o interface{}) (interface{}, error) {
	i := co.MustGet("daryl_id").(string)
	r := protodef.RefuseSessionSliceRequest{DarylIdentifier: i, Index: o.(*protodef.SessionSliceIndex)}
	return d.RefuseSessionSlice(context.Background(), &r)
}

var cmds = map[string]darylCommand{
	"message": &userMessageCommand{},
	"habit":   &addHabitCommand{},
	"session": &startWorkSessionCommand{},
	"cancel":  &cancelWorkSessionCommand{},
	"refuse":  &refuseSessionSliceCommand{},
}

func handleHTTPCommand(c *gin.Context) {
	cmdName := c.Param("command")
	cmd, ok := cmds[cmdName]
	if ok == false {
		c.JSON(400, gin.H{
			"status": "error", "error": "not found cmd",
		})
		c.Abort()
		return
	}
	url := c.MustGet("daryl_url").(string)
	o := cmd.Object()
	if o != nil {
		if err := c.BindJSON(&o); err != nil {
			log.Info(err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err})
			c.Abort()
			return
		}
	}
	d := openDarylConnection(url)
	resp, err := cmd.Execute(c, d, o)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"status": "ok",
		"resp":   resp,
	})
}

func handleCreateDaryl(c *gin.Context) {
	d := &protodef.Daryl{}
	if err := c.Bind(d); err != nil {
		c.JSON(500, gin.H{"status": "error", "error": err})
		c.Abort()
		return
	}

	f := openFarmConnection("localhost:8043")
	r, err := f.StartDaryl(context.Background(), &protodef.StartDarylRequest{Daryl: d})
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "error": err})
		c.Abort()
		return
	}

	t, err := newTokenForDaryl(r.Daryl)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "error": err})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"status": "ok",
		"daryl":  gin.H{"id": r.Daryl.Id},
		"token":  gin.H{"hash": t.Hash},
	})
}

func findDarylServer() func(string) (string, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{config.AppContext.String("etcd-url")},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	//defer cli.Close()

	return func(identifier string) (string, error) {
		resp, err := cli.Get(context.Background(), fmt.Sprintf("daryl_%s", identifier))
		if err != nil {
			return "", err
		}
		if len(resp.Kvs) != 0 {
			url := string(resp.Kvs[0].Value)
			return url, nil
		}
		return "", errors.New("Daryl not found")
	}
}

func setDarylServer() func(*gin.Context) {
	fds := findDarylServer()
	return func(c *gin.Context) {
		h := c.GetHeader(AUTH_TOKEN_HEADER)
		if h == "" {
			h = c.Param("token")
		}
		if h == "" {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "error", "error": errors.New("Access denied")})
			c.Abort()
			return
		}
		t, err := newTokenFromToken(h)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "error", "error": err})
			c.Abort()
			return
		}
		url, err := fds(t.Daryl.Id)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "error", "error": err})
			c.Abort()
			return
		}
		log.Infof("Daryl at %s", url)
		c.Set("daryl_url", url)
		c.Set("daryl_id", t.Daryl.Id)
	}
}
