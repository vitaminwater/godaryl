package daryl

import (
	log "github.com/sirupsen/logrus"
)

type noteMessageProcessor struct {
}

func (nmp *noteMessageProcessor) matches(r *UserMessageRequest) bool {
	return true
}

func (nmp *noteMessageProcessor) process(r *UserMessageRequest) {
	log.Info("noteMessageProcessor.process ", r)
}

func newNoteMessageProcessor() *noteMessageProcessor {
	nmp := &noteMessageProcessor{}
	return nmp
}
