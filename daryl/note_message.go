package daryl

import (
	log "github.com/sirupsen/logrus"
)

type noteMessage struct {
	text string
}

type noteMessageProcessor struct {
}

func (nmp *noteMessageProcessor) matches(r *UserMessageRequest) bool {
	return true
}

func (nmp *noteMessageProcessor) process(mr *messageRouter, r *UserMessageRequest) {
	mr.s.pub(noteMessage{r.Text}, NOTE_LOG_TOPIC)
	log.Info("noteMessageProcessor.process ", r)
}

func newNoteMessageProcessor() *noteMessageProcessor {
	nmp := &noteMessageProcessor{}
	return nmp
}
