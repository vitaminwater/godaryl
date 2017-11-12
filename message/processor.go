package message

import (
	"github.com/vitaminwater/daryl/daryl"
)

type messageProcessor struct {
	d  *daryl.Daryl
	mr *messageRouter
}

func (mp *messageProcessor) SetDaryl(d *daryl.Daryl) {
	mp.d = d
	mp.mr = newMessageRouter(d)
}

func (mp *messageProcessor) UserMessage(r *daryl.UserMessageRequest) (*daryl.UserMessageResponse, error) {
	mp.d.Pub(r, daryl.USER_MESSAGE_TOPIC)
	return &daryl.UserMessageResponse{}, nil
}

func NewMessageProcessor() *messageProcessor {
	mp := &messageProcessor{}
	return mp
}
