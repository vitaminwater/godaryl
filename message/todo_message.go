package message

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

type todoMessageProcessor struct {
	mp *messageProcessor
}

func (tmp *todoMessageProcessor) process(m *message) {
	if strings.HasPrefix(strings.ToLower(m.Text), "todo") {
		return
	}
	tmp.mp.d.Pub(m, TODO_LOG_TOPIC)
	log.Info("todoMessageProcessor.process")
}

func newTodoMessageProcessor(mp *messageProcessor) *todoMessageProcessor {
	tmp := &todoMessageProcessor{mp}
	return tmp
}
