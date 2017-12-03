package message

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/model"
)

type todoMessageProcessor struct {
	mp *messageProcessor
}

func (tmp *todoMessageProcessor) process(m model.Message) {
	if strings.HasPrefix(strings.ToLower(m.Text), "todo") == false {
		return
	}
	m.Attrs["todo"] = map[string]interface{}{
		"text": m.Text[5:],
		"done": false,
	}
	tmp.mp.d.Pub(m, TODO_LOG_TOPIC)
	log.Info("todoMessageProcessor.process")
}

func newTodoMessageProcessor(mp *messageProcessor) *todoMessageProcessor {
	tmp := &todoMessageProcessor{mp}
	return tmp
}
