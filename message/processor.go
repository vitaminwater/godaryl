package message

import (
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/model"
	"github.com/vitaminwater/daryl/protodef"
)

type messageTypeProcessor interface {
	process(model.Message)
}

type messageProcessor struct {
	d          *daryl.Daryl
	processors []messageTypeProcessor
}

func (mp *messageProcessor) SetDaryl(d *daryl.Daryl) {
	mp.d = d
}

func (mp *messageProcessor) UserMessage(r *protodef.UserMessageRequest) (*protodef.UserMessageResponse, error) {
	m := model.NewMessageFromProtodef(mp.d.D, r.Message)
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
