package daryl

import (
	log "github.com/sirupsen/logrus"
	"regexp"
)

type linkMessageProcessor struct {
}

var linkRegexp = regexp.MustCompile(`^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)

func (lmp *linkMessageProcessor) matches(r *UserMessageRequest) bool {
	return linkRegexp.Match([]byte(r.Text))
}

func (lmp *linkMessageProcessor) process(r *UserMessageRequest) {
	log.Info("linkMessageProcessor.process")
}

func newLinkMessageProcessor() *linkMessageProcessor {
	lmp := &linkMessageProcessor{}
	return lmp
}
