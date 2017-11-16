package daryl

import (
	"github.com/cskr/pubsub"
	"github.com/vitaminwater/daryl/protodef"
)

/*
	Weight calculation:
	- short period -> higher weight
	- long period -> higher weight
	- avg period -> no weight change
*/

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
	d.pubsub.Pub(TopicMessage{msgType, msg}, ALL_TOPIC)
	d.pubsub.Pub(TopicMessage{msgType, msg}, msgType)
	for _, topic := range topics {
		d.pubsub.Pub(TopicMessage{msgType, msg}, topic)
	}
}

func NewDaryl(identifier string, messageProcessor MessageProcessor, habitProcessor HabitProcessor, sessionProcessor SessionProcessor) *Daryl {
	d := &Daryl{identifier, pubsub.New(10), messageProcessor, habitProcessor, sessionProcessor}
	messageProcessor.SetDaryl(d)
	habitProcessor.SetDaryl(d)
	sessionProcessor.SetDaryl(d)
	return d
}
