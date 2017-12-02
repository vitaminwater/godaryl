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
	"github.com/vitaminwater/daryl/model"
	"github.com/vitaminwater/daryl/protodef"
)

type MessageProcessor interface {
	SetDaryl(*Daryl)

	/* RPC */
	UserMessage(*protodef.UserMessageRequest) (*protodef.UserMessageResponse, error)
}

type HabitProcessor interface {
	SetDaryl(*Daryl)

	/* RPC */
	AddHabit(*protodef.AddHabitRequest) (*protodef.AddHabitResponse, error)

	/* API */
	GetDueHabits() []model.Habit
	GetWeight(model.Habit) int
}

type SessionProcessor interface {
	SetDaryl(*Daryl)

	/* RPC */
	StartWorkSession(*protodef.StartWorkSessionRequest) (*protodef.StartWorkSessionResponse, error)
	CancelWorkSession(*protodef.CancelWorkSessionRequest) (*protodef.CancelWorkSessionResponse, error)
	RefuseSessionSlice(*protodef.RefuseSessionSliceRequest) (*protodef.RefuseSessionSliceResponse, error)
}

type Daryl struct {
	D      model.Daryl
	pubsub *pubsub.PubSub

	MessageProcessor MessageProcessor
	HabitProcessor   HabitProcessor
	SessionProcessor SessionProcessor
}

func (d *Daryl) Sub(topics ...string) chan interface{} {
	return d.pubsub.Sub(topics...)
}

func (d *Daryl) Pub(msg interface{}, msgType string, topics ...string) {
	m := TopicMessage{msgType, msg}
	j, err := json.Marshal(m)
	if err != nil {
		log.Info(err)
	}
	c := kv.Pool.Get()

	d.pubsub.Pub(m, ALL_TOPIC)
	d.pubsub.Pub(m, msgType)
	if len(j) != 0 {
		c.Do("PUBLISH", fmt.Sprintf("daryl.%s.%s", d.D.Id, msgType), string(j))
	}
	for _, topic := range topics {
		d.pubsub.Pub(m, topic)
		if len(j) != 0 {
			c.Do("PUBLISH", fmt.Sprintf("daryl.%s.%s", d.D.Id, topic), string(j))
		}
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func NewDaryl(da model.Daryl, messageProcessor MessageProcessor, habitProcessor HabitProcessor, sessionProcessor SessionProcessor) *Daryl {
	d := &Daryl{
		D:                da,
		pubsub:           pubsub.New(10),
		MessageProcessor: messageProcessor,
		HabitProcessor:   habitProcessor,
		SessionProcessor: sessionProcessor,
	}
	messageProcessor.SetDaryl(d)
	habitProcessor.SetDaryl(d)
	sessionProcessor.SetDaryl(d)
	go distributed.Beacon(fmt.Sprintf("daryl_%s", d.D.Id), config.AppContext.String("advertized-url"))
	return d
}
