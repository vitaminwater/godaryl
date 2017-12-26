package message

import (
	"github.com/vitaminwater/daryl/model"
)

type threadCommand interface {
	execute(t *thread) error
}

/**
 * Add conversation
 */

type addConversationCommand struct {
	c conversation
}

func (c addConversationCommand) execute(t *thread) error {
	t.cs = append(t.cs, c.c)
	t.updateCurrentConversation()
	return nil
}

/**
 * Stop current conversation
 */

type stopCurrentConversationCommand struct {
}

func (c stopCurrentConversationCommand) execute(t *thread) error {
	t.c = nil
	t.updateCurrentConversation()
	return nil
}

/**
 * Push user message
 */

type pushUserMessageCommand struct {
	m model.Message
}

func (c pushUserMessageCommand) execute(t *thread) error {
	if t.c == nil {
		return nil
	}
	keep, err := t.c.pushUserMessage(c.m)
	if !keep {
		t.c = nil
		t.updateCurrentConversation()
	}
	if err != nil {
		return err
	}
	return nil
}

/**
 * Push user message
 */

type getUserMessagesCommand struct {
	from int32
	to   int32
	r    chan []model.Message
}

func (c getUserMessagesCommand) execute(t *thread) error {
	c.r <- []model.Message{}
	return nil
}
