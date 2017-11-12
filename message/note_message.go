package message

import (
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/protodef"
)

type noteMessage struct {
	text string
}

type noteMessageProcessor struct {
}

func (nmp *noteMessageProcessor) matches(r *protodef.UserMessageRequest) bool {
	return true
}

func (nmp *noteMessageProcessor) process(mr *messageRouter, r *protodef.UserMessageRequest) {
	mr.d.Pub(noteMessage{r.Text}, NOTE_LOG_TOPIC)
	log.Info("noteMessageProcessor.process ", r)
}

func newNoteMessageProcessor() *noteMessageProcessor {
	nmp := &noteMessageProcessor{}
	return nmp
}
