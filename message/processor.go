package message

import (
	"fmt"

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
	threads    []thread
}

func (mp *messageProcessor) SetDaryl(d *daryl.Daryl) {
	mp.d = d
}

func (mp *messageProcessor) UserMessage(r *protodef.UserMessageRequest) (*protodef.UserMessageResponse, error) {
	m, err := model.NewMessageFromProtodef(mp.d.D, r.Message)
	if err != nil {
		return nil, err
	}

	err = m.Insert()
	if err != nil {
		return nil, err
	}

	for _, processor := range mp.processors {
		processor.process(m)
	}

	err = m.Update()
	if err != nil {
		return nil, err
	}

	mp.d.Pub(m, daryl.USER_MESSAGE_TOPIC, fmt.Sprintf("%s.%s", daryl.USER_MESSAGE_TOPIC, m.HabitId.String))

	for _, t := range mp.threads {
		t.pushUserMessage(m)
	}

	mm, err := m.ToProtodef()
	if err != nil {
		return nil, err
	}
	return &protodef.UserMessageResponse{Message: mm}, nil
}

func NewMessageProcessor() *messageProcessor {
	mp := &messageProcessor{}
	mp.processors = []messageTypeProcessor{
		newLinkMessageProcessor(mp),
		newTodoMessageProcessor(mp),
	}
	return mp
}
