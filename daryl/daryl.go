package daryl

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cskr/pubsub"
	"github.com/labstack/gommon/log"
	"github.com/vitaminwater/daryl/config"
	"github.com/vitaminwater/daryl/distributed"
	"github.com/vitaminwater/daryl/kv"
	"github.com/vitaminwater/daryl/protodef"
)

type Habit interface {
	GetHabit() protodef.Habit
	GetWeight() uint32
}

type SessionSlice interface {
	GetSessionSlice() protodef.SessionSlice
}

type Session interface {
	GetSession() protodef.Session
	GetSessionSlices() []SessionSlice
}

type MessageProcessor interface {
	SetDaryl(*Daryl)
	UserMessage(*protodef.UserMessageRequest) (*protodef.UserMessageResponse, error)
}

type HabitProcessor interface {
	SetDaryl(*Daryl)

	/* RPC */
	AddHabit(*protodef.AddHabitRequest) (*protodef.AddHabitResponse, error)

	/* API */
	GetDueHabits() []Habit
}

type SessionProcessor interface {
	SetDaryl(*Daryl)
	StartWorkSession(*protodef.StartWorkSessionRequest) (*protodef.StartWorkSessionResponse, error)
	CancelWorkSession(*protodef.CancelWorkSessionRequest) (*protodef.CancelWorkSessionResponse, error)
	RefuseSessionSlice(*protodef.RefuseSessionSliceRequest) (*protodef.RefuseSessionSliceResponse, error)
}

type Daryl struct {
	identifier string
	pubsub     *pubsub.PubSub

	MessageProcessor MessageProcessor
	HabitProcessor   HabitProcessor
	SessionProcessor SessionProcessor
}

func (d *Daryl) Sub(topics ...string) chan interface{} {
	return d.pubsub.Sub(topics...)
}

func (d *Daryl) Pub(msg interface{}, msgType string, topics ...string) {
	m := TopicMessage{msgType, msg}
	d.pubsub.Pub(m, ALL_TOPIC)
	d.pubsub.Pub(m, msgType)
	for _, topic := range topics {
		d.pubsub.Pub(m, topic)
	}

	c := kv.Pool.Get()

	j, err := json.Marshal(m)
	if err != nil {
		log.Info(err)
	}
	c.Do("PUBLISH", fmt.Sprintf("daryl.%s", d.identifier), string(j))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func NewDaryl(identifier string, messageProcessor MessageProcessor, habitProcessor HabitProcessor, sessionProcessor SessionProcessor) *Daryl {
	d := &Daryl{identifier, pubsub.New(10), messageProcessor, habitProcessor, sessionProcessor}
	messageProcessor.SetDaryl(d)
	habitProcessor.SetDaryl(d)
	sessionProcessor.SetDaryl(d)
	go distributed.Beacon(fmt.Sprintf("daryl_%s", d.identifier), config.AppContext.String("advertized-url"))
	return d
}
