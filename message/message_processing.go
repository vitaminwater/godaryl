package message

import (
	"github.com/vitaminwater/daryl/daryl"
)

type messageTypeProcessor interface {
	matches(*daryl.UserMessageRequest) bool
	process(*messageRouter, *daryl.UserMessageRequest)
}

type messageRouter struct {
	d          *daryl.Daryl
	c          chan interface{}
	processors []messageTypeProcessor
}

func messageRouterProcess(mr *messageRouter) {
	for msg := range mr.c {
		tm := msg.(daryl.TopicMessage)
		r := tm.Msg.(*daryl.UserMessageRequest)
		for _, processor := range mr.processors {
			if processor.matches(r) {
				processor.process(mr, r)
				break
			}
		}
	}
}

func newMessageRouter(d *daryl.Daryl) *messageRouter {
	mr := &messageRouter{d: d}
	mr.c = d.Sub(
		daryl.USER_MESSAGE_TOPIC,
	)
	mr.processors = []messageTypeProcessor{
		newTodoMessageProcessor(),
		newLinkMessageProcessor(),
		newNoteMessageProcessor(),
	}
	go messageRouterProcess(mr)
	return mr
}
