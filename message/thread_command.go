package message

import (
	"github.com/vitaminwater/daryl/model"
)

type threadCommand interface {
	execute(t *thread) error
}

type addConversationCommand struct {
	c conversation
}

func (c addConversationCommand) execute(t *thread) error {
	t.cs = append(t.cs, c.c)
	t.updateCurrentConversation()
	return nil
}

type stopCurrentConversationCommand struct {
}

func (c stopCurrentConversationCommand) execute(t *thread) error {
	t.c = nil
	t.updateCurrentConversation()
	return nil
}

type pushUserMessageCommand struct {
	m model.Message
}

func (c pushUserMessageCommand) execute(t *thread) error {
	if t.c == nil {
		return nil
	}
	keep, err := t.c.pushUserMessage(c.m)
	if err != nil {
		return err
	}
	if !keep {
		t.c = nil
		t.updateCurrentConversation()
	}
	return nil
}
