package message

import (
	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/protodef"
	"time"
)

type message struct {
	protodef.Message
	At time.Time `db:"at"`
}

func newMessage(msg *protodef.Message) *message {
	at, err := ptypes.Timestamp(msg.At)
	if err != nil {
		log.Warning(err)
	}
	m := &message{*msg, at}
	return m
}

type messageTypeProcessor interface {
	process(*message)
}

type messageProcessor struct {
	d          *daryl.Daryl
	processors []messageTypeProcessor
}

func (mp *messageProcessor) SetDaryl(d *daryl.Daryl) {
	mp.d = d
}

func (mp *messageProcessor) UserMessage(r *protodef.UserMessageRequest) (*protodef.UserMessageResponse, error) {
	m := newMessage(r.Message)
	mp.d.Pub(m, daryl.USER_MESSAGE_TOPIC)
	for _, processor := range mp.processors {
		processor.process(m)
	}
	return &protodef.UserMessageResponse{}, nil
}

func NewMessageProcessor() *messageProcessor {
	mp := &messageProcessor{}
	mp.processors = []messageTypeProcessor{
		newLinkMessageProcessor(mp),
		newTodoMessageProcessor(mp),
	}
	return mp
}
