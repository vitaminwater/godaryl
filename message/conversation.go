package message

import "github.com/vitaminwater/daryl/model"

type conversation interface {
	isReady() bool
	addMessage(model.Message) (bool, error)
}

type HabitConversation struct {
}

func (hc HabitConversation) isReady() bool {
	return false
}

func (hc HabitConversation) addMessage() (bool, error) {
	return false, nil
}
