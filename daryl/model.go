package daryl

import (
	"github.com/vitaminwater/daryl/model"
	"github.com/vitaminwater/daryl/protodef"
)

type Habit interface {
	GetHabit() model.Habit
}

type MessageProcessor interface {
	SetDaryl(*Daryl)

	/* RPC */
	UserMessage(*protodef.UserMessageRequest) (*protodef.UserMessageResponse, error)
}

type HabitProcessor interface {
	SetDaryl(*Daryl)

	/* RPC */
	AddHabit(*protodef.AddHabitRequest) (*protodef.AddHabitResponse, error)

	/* API */
	GetHabit(string) (model.Habit, error)
	GetDueHabits() []model.Habit
	GetWeight(model.Habit) int
}

type TriggerProcessor interface {
	SetDaryl(*Daryl)

	/* RPC */
	AddTrigger(*protodef.AddTriggerRequest) (*protodef.AddTriggerResponse, error)
}

type SessionProcessor interface {
	SetDaryl(*Daryl)

	/* RPC */
	StartWorkSession(*protodef.StartWorkSessionRequest) (*protodef.StartWorkSessionResponse, error)
	CancelWorkSession(*protodef.CancelWorkSessionRequest) (*protodef.CancelWorkSessionResponse, error)
	RefuseSessionSlice(*protodef.RefuseSessionSliceRequest) (*protodef.RefuseSessionSliceResponse, error)
}
