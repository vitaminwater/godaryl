package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitaminwater/daryl/protodef"
)

type darylCommand interface {
	Name() string
	Object() interface{}
	Execute(*gin.Context, protodef.DarylServiceClient, interface{}) (interface{}, error)
}

/**
 * Message
 */

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
	r := protodef.UserMessageRequest{DarylIdentifier: i, Message: o.(*protodef.Message)}
	return d.UserMessage(context.Background(), &r)
}

/**
 * GetMessages
 */

type getUserMessagesCommand struct {
}

func (c *getUserMessagesCommand) Name() string {
	return "getmessages"
}

func (c *getUserMessagesCommand) Object() interface{} {
	return &protodef.GetUserMessagesRequest{}
}

func (c *getUserMessagesCommand) Execute(co *gin.Context, d protodef.DarylServiceClient, o interface{}) (interface{}, error) {
	i := co.MustGet("daryl_id").(string)
	r := o.(*protodef.GetUserMessagesRequest)
	r.DarylIdentifier = i
	return d.GetUserMessages(context.Background(), r)
}

/**
 * Habit
 */

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
	r := protodef.AddHabitRequest{DarylIdentifier: i, Habit: (o.(*protodef.Habit))}
	return d.AddHabit(context.Background(), &r)
}

/**
 * GetHabit
 */

type getHabitsCommand struct {
}

func (c *getHabitsCommand) Name() string {
	return "habit"
}

func (c *getHabitsCommand) Object() interface{} {
	return &protodef.GetHabitsRequest{}
}

func (c *getHabitsCommand) Execute(co *gin.Context, d protodef.DarylServiceClient, o interface{}) (interface{}, error) {
	i := co.MustGet("daryl_id").(string)
	r := o.(*protodef.GetHabitsRequest)
	r.DarylIdentifier = i
	return d.GetHabits(context.Background(), r)
}

/**
 * Trigger
 */

type addTriggerCommand struct {
}

func (c *addTriggerCommand) Name() string {
	return "trigger"
}

func (c *addTriggerCommand) Object() interface{} {
	return &protodef.Trigger{}
}

func (c *addTriggerCommand) Execute(co *gin.Context, d protodef.DarylServiceClient, o interface{}) (interface{}, error) {
	i := co.MustGet("daryl_id").(string)
	r := protodef.AddTriggerRequest{DarylIdentifier: i, Trigger: (o.(*protodef.Trigger))}
	return d.AddTrigger(context.Background(), &r)
}

/**
 * Work session
 */

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

/**
 * Get Work session
 */

type getWorkSessionCommand struct {
}

func (c *getWorkSessionCommand) Name() string {
	return "session"
}

func (c *getWorkSessionCommand) Object() interface{} {
	return &protodef.GetWorkSessionRequest{}
}

func (c *getWorkSessionCommand) Execute(co *gin.Context, d protodef.DarylServiceClient, o interface{}) (interface{}, error) {
	i := co.MustGet("daryl_id").(string)
	r := o.(*protodef.GetWorkSessionRequest)
	r.DarylIdentifier = i
	return d.GetWorkSession(context.Background(), r)
}

/**
 * Cancel work session
 */

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

/**
 * Refuse session slice
 */

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

/**
 * Get
 */

type getCommand struct {
}

func (c *getCommand) Name() string {
	return "get"
}

func (c *getCommand) Object() interface{} {
	return nil
}

func (c *getCommand) Execute(co *gin.Context, d protodef.DarylServiceClient, o interface{}) (interface{}, error) {
	i := co.MustGet("daryl_id").(string)
	r := protodef.GetRequest{DarylIdentifier: i}
	return d.Get(context.Background(), &r)
}

/**
 * handleHTTPCommand
 */

var posts = map[string]darylCommand{
	"message": &userMessageCommand{},
	"habit":   &addHabitCommand{},
	"trigger": &addTriggerCommand{},
	"session": &startWorkSessionCommand{},
	"cancel":  &cancelWorkSessionCommand{},
	"refuse":  &refuseSessionSliceCommand{},
	"get":     &getCommand{},
}

var gets = map[string]darylCommand{
	"message": &getUserMessagesCommand{},
	"habit":   &getHabitsCommand{},
	"session": &getWorkSessionCommand{},
	"get":     &getCommand{},
}

func handleHTTPCommand(cmds map[string]darylCommand) func(c *gin.Context) {
	return func(c *gin.Context) {
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
