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
)

type Daryl struct {
	D      model.Daryl
	pubsub *pubsub.PubSub

	MessageProcessor MessageProcessor
	HabitProcessor   HabitProcessor
	TriggerProcessor TriggerProcessor
	SessionProcessor SessionProcessor
}

func (d *Daryl) Sub(topics ...string) chan interface{} {
	return d.pubsub.Sub(topics...)
}

func (d *Daryl) Pub(msg interface{}, msgType string, topics ...string) {
	m := TopicMessage{ALL_TOPIC, msg}
	c := kv.Pool.Get()

	d.pubsub.Pub(m, ALL_TOPIC)
	m.Topic = msgType
	d.pubsub.Pub(m, msgType)

	p, err := model.ToProtodef(msg)
	if err != nil {
		log.Info(err)
	} else {
		j, err := json.Marshal(TopicMessage{fmt.Sprintf("daryl.%s.%s", d.D.Id, msgType), p})
		if err != nil {
			log.Info(err)
		}
		c.Do("PUBLISH", fmt.Sprintf("daryl.%s.%s", d.D.Id, msgType), string(j))
	}
	for _, topic := range topics {
		m.Topic = topic
		d.pubsub.Pub(m, topic)

		p, err := model.ToProtodef(msg)
		if err != nil {
			log.Info(err)
		} else {
			j, err := json.Marshal(TopicMessage{fmt.Sprintf("daryl.%s.%s", d.D.Id, topic), p})
			if err != nil {
				log.Info(err)
			}
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

func NewDaryl(da model.Daryl, messageProcessor MessageProcessor, habitProcessor HabitProcessor, triggerProcessor TriggerProcessor, sessionProcessor SessionProcessor) *Daryl {
	d := &Daryl{
		D:                da,
		pubsub:           pubsub.New(10),
		MessageProcessor: messageProcessor,
		HabitProcessor:   habitProcessor,
		TriggerProcessor: triggerProcessor,
		SessionProcessor: sessionProcessor,
	}
	habitProcessor.SetDaryl(d)
	triggerProcessor.SetDaryl(d)
	sessionProcessor.SetDaryl(d)
	messageProcessor.SetDaryl(d)
	go distributed.Beacon(fmt.Sprintf("daryl_%s", d.D.Id), config.AppContext.String("advertized-url"))
	return d
}
