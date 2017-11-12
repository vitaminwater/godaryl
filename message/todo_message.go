package message

import (
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/protodef"
	"strings"
)

type todoMessage struct {
	text string
}

type todoMessageProcessor struct {
}

func (lmp *todoMessageProcessor) matches(r *protodef.UserMessageRequest) bool {
	return strings.HasPrefix(strings.ToLower(r.Text), "todo")
}

func (lmp *todoMessageProcessor) process(mr *messageRouter, r *protodef.UserMessageRequest) {
	mr.d.Pub(todoMessage{r.Text}, TODO_LOG_TOPIC)
	log.Info("todoMessageProcessor.process")
}

func newTodoMessageProcessor() *todoMessageProcessor {
	lmp := &todoMessageProcessor{}
	return lmp
}
