package message

import (
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/protodef"
)

type messageProcessor struct {
	d  *daryl.Daryl
	mr *messageRouter
}

func (mp *messageProcessor) SetDaryl(d *daryl.Daryl) {
	mp.d = d
	mp.mr = newMessageRouter(d)
}

func (mp *messageProcessor) UserMessage(r *protodef.UserMessageRequest) (*protodef.UserMessageResponse, error) {
	mp.d.Pub(r, daryl.USER_MESSAGE_TOPIC)
	return &protodef.UserMessageResponse{}, nil
}

func NewMessageProcessor() *messageProcessor {
	mp := &messageProcessor{}
	return mp
}