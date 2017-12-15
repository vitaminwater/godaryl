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
	c  *conversation

	cmd chan threadCommand
}

func (t thread) stopCurrentConversation() {
	t.cmd <- stopCurrentConversationCommand{}
}

func (t thread) addConversation(c conversation) {
	t.cmd <- addConversationCommand{c: c}
}

func (t thread) pushUserMessage(m model.Message) {
	t.cmd <- pushUserMessageCommand{m: m}
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

func newThread(id string, d *daryl.Daryl) (thread, error) {
	t := thread{id: id, d: d, cs: []conversation{}}

	go threadProcess(t)
	return t, nil
}
