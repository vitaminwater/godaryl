package daryl

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

type todoMessageProcessor struct {
}

func (lmp *todoMessageProcessor) matches(r *UserMessageRequest) bool {
	return strings.HasPrefix(strings.ToLower(r.Text), "todo")
}

func (lmp *todoMessageProcessor) process(r *UserMessageRequest) {
	log.Info("todoMessageProcessor.process")
}

func newTodoMessageProcessor() *todoMessageProcessor {
	lmp := &todoMessageProcessor{}
	return lmp
}
