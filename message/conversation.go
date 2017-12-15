package message

import (
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/model"
)

type conversation interface {
	isReady() bool
	priority() int
	pushUserMessage(model.Message) (bool, error)
}

/**
 * Habit conversation
 */

type habitConversation struct {
	h daryl.Habit
}

func (hc habitConversation) priority() int {
	return 0
}

func (hc habitConversation) isReady() bool {
	return true
}

func (hc habitConversation) pushUserMessage(m model.Message) (bool, error) {
	log.Info("habitConversation", m)
	return true, nil
}

/**
 * Daryl conversation
 */

type darylConversation struct {
	d *daryl.Daryl
}

func (hc darylConversation) priority() int {
	return 0
}

func (dc darylConversation) isReady() bool {
	return true
}

func (dc darylConversation) pushUserMessage(m model.Message) (bool, error) {
	log.Info("darylConversation", m)
	return true, nil
}
