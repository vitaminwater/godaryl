package daryl

import (
	"github.com/cskr/pubsub"
)

type Daryl struct {
	identifier string
	pubsub     *pubsub.PubSub
}

func (d *Daryl) Sub(topic ...string) chan interface{} {
	return d.pubsub.Sub(topic...)
}

func (d *Daryl) Pub(msg interface{}, topic string) {
	d.pubsub.Pub(TopicMessage{topic, msg}, topic)
}

func NewDaryl(identifier string) *Daryl {
	d := &Daryl{identifier, pubsub.New(10)}
	return d
}
