package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitaminwater/daryl/model"
	"github.com/vitaminwater/daryl/protodef"
)

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
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err})
			c.Abort()
			return
		}
	}
	d, cl := protodef.OpenDarylConnection(url)
	defer cl()
	resp, err := cmd.Execute(c, d, o)
	if err != nil {
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

	f, cl := protodef.OpenFarmConnection("localhost:8043")
	defer cl()
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

func handleCreateDarylToken(c *gin.Context) {
	d := &protodef.Daryl{}
	if err := c.Bind(d); err != nil {
		c.JSON(500, gin.H{"status": "error", "error": err})
		c.Abort()
		return
	}

	da, err := model.NewDarylFromProtodef(d)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "error": err})
		c.Abort()
		return
	}

	if err := da.GetFromNameAndPassword(); err != nil {
		c.JSON(500, gin.H{"status": "error", "error": err})
		c.Abort()
		return
	}

	d, err = da.ToProtodef()
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "error": err})
		c.Abort()
		return
	}
	t, err := newTokenForDaryl(d)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "error": err})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"status": "ok",
		"token":  gin.H{"hash": t.Hash},
	})
}
