package message

import (
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/protodef"
	"regexp"
)

type linkMessage struct {
	text string
}

type linkMessageProcessor struct {
}

var linkRegexp = regexp.MustCompile(`^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)

func (lmp *linkMessageProcessor) matches(r *protodef.UserMessageRequest) bool {
	return linkRegexp.Match([]byte(r.Message.Text))
}

func (lmp *linkMessageProcessor) process(mr *messageRouter, r *protodef.UserMessageRequest) {
	mr.d.Pub(linkMessage{r.Message.Text}, LINK_LOG_TOPIC)
	log.Info("linkMessageProcessor.process")
}

func newLinkMessageProcessor() *linkMessageProcessor {
	lmp := &linkMessageProcessor{}
	return lmp
}
