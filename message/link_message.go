package message

import (
	"regexp"

	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/model"
)

type linkMessageProcessor struct {
	mp *messageProcessor
}

var linkRegexp = regexp.MustCompile(`(https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*))`)

func (lmp *linkMessageProcessor) process(m model.Message) {
	matches := linkRegexp.FindAllStringSubmatch(m.Text, 10)

	if len(matches) == 0 {
		return
	}

	l := []string{}
	for _, match := range matches {
		l = append(l, match[0])
	}

	m.Attrs["link"] = l
	lmp.mp.d.Pub(m, LINK_LOG_TOPIC)
	log.Info("linkMessageProcessor.process")
}

func newLinkMessageProcessor(mp *messageProcessor) *linkMessageProcessor {
	lmp := &linkMessageProcessor{mp}
	return lmp
}
