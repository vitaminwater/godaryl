package daryl

import (
	"github.com/cskr/pubsub"
)

type MessageProcessor interface {
	SetDaryl(*Daryl)
	UserMessage(*UserMessageRequest) (*UserMessageResponse, error)
}

type HabitProcessor interface {
	SetDaryl(*Daryl)
	AddHabit(*AddHabitRequest) (*AddHabitResponse, error)
}

type SessionProcessor interface {
	SetDaryl(*Daryl)
	StartWorkSession(*StartWorkSessionRequest) (*StartWorkSessionResponse, error)
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
