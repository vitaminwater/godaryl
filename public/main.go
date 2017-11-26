package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/vitaminwater/daryl/protodef"

	"github.com/coreos/etcd/clientv3"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const AUTH_TOKEN_HEADER = "X-Daryl-Auth-Token"

type DarylCommand interface {
	Name() string
	Object() interface{}
	Execute(protodef.DarylClient, interface{}) (interface{}, error)
}

type UserMessageCommand struct {
}

func (c *UserMessageCommand) Name() string {
	return "message"
}

func (c *UserMessageCommand) Object() interface{} {
	return &protodef.Message{}
}

func (c *UserMessageCommand) Execute(d protodef.DarylClient, o interface{}) (interface{}, error) {
	um := protodef.UserMessageRequest{Identifier: "lol", Message: o.(*protodef.Message)}
	return d.UserMessage(context.Background(), &um)
}

type AddHabitCommand struct {
	protodef.Habit
	Deadline time.Time `json:"deadline"`
}

func (c *AddHabitCommand) Name() string {
	return "habit"
}

func (c *AddHabitCommand) Object() interface{} {
	return &AddHabitCommand{}
}

func (c *AddHabitCommand) Execute(d protodef.DarylClient, o interface{}) (interface{}, error) {
	ah := protodef.AddHabitRequest{Identifier: "lol", Habit: &(o.(*AddHabitCommand).Habit)}

	deadline, err := ptypes.TimestampProto(c.Deadline)
	if err != nil {
		log.Fatal(err)
	}

	ah.Habit.Deadline = deadline
	ah.Habit.LastDone = ptypes.TimestampNow()
	return d.AddHabit(context.Background(), &ah)
}

type StartWorkSessionCommand struct {
}

func (c *StartWorkSessionCommand) Name() string {
	return "session"
}

func (c *StartWorkSessionCommand) Object() interface{} {
	return &protodef.SessionConfig{}
}

func (c *StartWorkSessionCommand) Execute(d protodef.DarylClient, o interface{}) (interface{}, error) {
	r := protodef.StartWorkSessionRequest{Identifier: "lol", Config: o.(*protodef.SessionConfig)}
	return d.StartWorkSession(context.Background(), &r)
}

type CancelWorkSessionCommand struct {
}

func (c *CancelWorkSessionCommand) Name() string {
	return "cancel"
}

func (c *CancelWorkSessionCommand) Object() interface{} {
	return nil
}

func (c *CancelWorkSessionCommand) Execute(d protodef.DarylClient, o interface{}) (interface{}, error) {
	r := protodef.CancelWorkSessionRequest{Identifier: "lol"}
	return d.CancelWorkSession(context.Background(), &r)
}

type RefuseSessionSliceCommand struct {
}

func (c *RefuseSessionSliceCommand) Name() string {
	return "refuse"
}

func (c *RefuseSessionSliceCommand) Object() interface{} {
	return &protodef.SessionSliceIndex{}
}

func (c *RefuseSessionSliceCommand) Execute(d protodef.DarylClient, o interface{}) (interface{}, error) {
	r := protodef.RefuseSessionSliceRequest{Identifier: "lol", Index: o.(*protodef.SessionSliceIndex)}
	return d.RefuseSessionSlice(context.Background(), &r)
}

var cmds = map[string]DarylCommand{
	"message": &UserMessageCommand{},
	"habit":   &AddHabitCommand{},
	"session": &StartWorkSessionCommand{},
	"cancel":  &CancelWorkSessionCommand{},
	"refuse":  &RefuseSessionSliceCommand{},
}

func handleHTTPCommand(c *gin.Context) {
	cmdName := c.Param("command")
	cmd, ok := cmds[cmdName]
	if ok == false {
		c.JSON(400, gin.H{
			"status": "not found cmd",
		})
		return
	}
	url := c.MustGet("daryl_url").(string)
	o := cmd.Object()
	if o != nil {
		if err := c.BindJSON(&o); err != nil {
			log.Info(err)
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}
	}
	d := openDarylConnection(url)
	resp, err := cmd.Execute(d, o)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	log.Info(url, o, resp)
	c.JSON(200, gin.H{
		"status": "posted",
	})

}

func findDarylServer() func(string) (string, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	//defer cli.Close()

	return func(identifier string) (string, error) {
		resp, err := cli.Get(context.Background(), fmt.Sprintf("daryl_%s", identifier))
		if err != nil {
			log.Fatal(err)
		}
		if len(resp.Kvs) != 0 {
			url := string(resp.Kvs[0].Value)
			log.Info(url)
			return url, nil
		}
		return "", errors.New("Daryl not found")
	}
}

func setDarylServer() func(*gin.Context) {
	fds := findDarylServer()
	return func(c *gin.Context) {
		url, err := fds("lol")
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": err})
			c.Abort()
			return
		}
		log.Infof("Daryl at %s", url)
		c.Set("daryl_url", url)
	}
}

func main() {
	router := gin.Default()
	router.Use(setDarylServer())
	router.POST("/daryl/cmd/:command", handleHTTPCommand)
	router.Run()
}
