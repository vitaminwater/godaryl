package message

import (
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/model"
)

type thread struct {
	id string
	d  *daryl.Daryl
	cs []conversation
	c  conversation

	cmd chan threadCommand
}

/**
 * Commands wrappers
 */

func (t thread) stopCurrentConversation() {
	t.cmd <- stopCurrentConversationCommand{}
}

func (t thread) addConversation(c conversation) {
	t.cmd <- addConversationCommand{c: c}
}

func (t thread) pushUserMessage(m model.Message) {
	t.cmd <- pushUserMessageCommand{m: m}
}

func (t thread) getUserMessages(from, to int32) []model.Message {
	r := make(chan []model.Message)
	t.cmd <- getUserMessagesCommand{from, to, r}
	res := <-r
	return res
}

/**
 * internals
 */

func (t *thread) updateCurrentConversation() {
	var newCurrent = t.c
	for _, c := range t.cs {
		if newCurrent == nil || (c.isReady() && newCurrent.priority() < c.priority()) {
			newCurrent = c
		}
	}
	t.c = newCurrent
}

func threadProcess(t thread) {
	for {
		select {
		case c := <-t.cmd:
			if err := c.execute(&t); err != nil {
				log.Warning(err)
			}
		}
	}
}

func newThread(id string, d *daryl.Daryl, cs []conversation) (thread, error) {
	t := thread{id: id, d: d, cs: cs, c: nil, cmd: make(chan threadCommand, 10)}

	t.updateCurrentConversation()
	go threadProcess(t)
	return t, nil
}
