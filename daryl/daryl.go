package daryl

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"
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
	CancelWorkSession(*protodef.CancelWorkSessionRequest) (*protodef.CancelWorkSessionResponse, error)
	RefuseWorkSession(*protodef.RefuseWorkSessionRequest) (*protodef.RefuseWorkSessionResponse, error)
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

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func advertiseEtcd(d *Daryl) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	for {
		gr, err := cli.Grant(context.TODO(), 2)
		if err != nil {
			log.Fatal(err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_, err = cli.Put(ctx, fmt.Sprintf("daryl_%s", d.identifier), getEnv("PUBLIC_URL", "http://localhost:8080"), clientv3.WithLease(gr.ID))
		cancel()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func NewDaryl(identifier string, messageProcessor MessageProcessor, habitProcessor HabitProcessor, sessionProcessor SessionProcessor) *Daryl {
	d := &Daryl{identifier, pubsub.New(10), messageProcessor, habitProcessor, sessionProcessor}
	messageProcessor.SetDaryl(d)
	habitProcessor.SetDaryl(d)
	sessionProcessor.SetDaryl(d)
	go advertiseEtcd(d)
	return d
}
