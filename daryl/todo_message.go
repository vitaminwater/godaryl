package daryl

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

type todoMessage struct {
	text string
}

type todoMessageProcessor struct {
}

func (lmp *todoMessageProcessor) matches(r *UserMessageRequest) bool {
	return strings.HasPrefix(strings.ToLower(r.Text), "todo")
}

func (lmp *todoMessageProcessor) process(mr *messageRouter, r *UserMessageRequest) {
	mr.s.pub(todoMessage{r.Text}, TODO_LOG_TOPIC)
	log.Info("todoMessageProcessor.process")
}

func newTodoMessageProcessor() *todoMessageProcessor {
	lmp := &todoMessageProcessor{}
	return lmp
}
