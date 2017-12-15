package message

import "github.com/vitaminwater/daryl/model"

type threadCommand interface {
	execute(t *thread) error
}

type addConversationCommand struct {
	c conversation
}

func (c addConversationCommand) execute(t *thread) error {
	t.cs = append(t.cs, c.c)
	return nil
}

type stopCurrentConversationCommand struct {
}

func (c stopCurrentConversationCommand) execute(t *thread) error {
	t.c = nil
	return nil
}

type pushUserMessageCommand struct {
	m model.Message
}

func (c pushUserMessageCommand) execute(t *thread) error {
	return nil
}
