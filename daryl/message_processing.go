package daryl

import ()

type messageProcessor interface {
	matches(*UserMessageRequest) bool
	process(*UserMessageRequest)
}

type messageRouter struct {
	c          chan interface{}
	processors []messageProcessor
}

func messageRouterProcess(mr *messageRouter) {
	for msg := range mr.c {
		tm := msg.(topicMessage)
		r := tm.msg.(*UserMessageRequest)
		for _, processor := range mr.processors {
			if processor.matches(r) {
				processor.process(r)
				break
			}
		}
	}
}

func newMessageRouter(s *darylServer) *messageRouter {
	mr := &messageRouter{}
	mr.c = s.pubsub.Sub(
		USER_MESSAGE_TOPIC,
	)
	mr.processors = []messageProcessor{
		newTodoMessageProcessor(),
		newLinkMessageProcessor(),
		newNoteMessageProcessor(),
	}
	go messageRouterProcess(mr)
	return mr
}
