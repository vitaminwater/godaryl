package daryl

import (
	"github.com/vitaminwater/daryl/model"
	"github.com/vitaminwater/daryl/protodef"
)

type Trigger interface {
}

type Habit interface {
	GetHabit() model.Habit
	Trigger(Trigger)
	GetWeight() int
}

type MessageProcessor interface {
	SetDaryl(*Daryl)

	/* RPC */
	UserMessage(*protodef.UserMessageRequest) (*protodef.UserMessageResponse, error)
	GetUserMessages(*protodef.GetUserMessagesRequest) (*protodef.GetUserMessagesResponse, error)
}

type HabitProcessor interface {
	SetDaryl(*Daryl)

	/* RPC */
	AddHabit(*protodef.AddHabitRequest) (*protodef.AddHabitResponse, error)
	GetHabits(*protodef.GetHabitsRequest) (*protodef.GetHabitsResponse, error)

	/* API */
	GetHabit(string) (Habit, error)
	GetAllHabits() ([]Habit, error)
	GetDueHabits() []Habit
}

type TriggerProcessor interface {
	SetDaryl(*Daryl)

	/* RPC */
	AddTrigger(*protodef.AddTriggerRequest) (*protodef.AddTriggerResponse, error)
	IncomingMessage(m *protodef.IncomingTriggerMessageRequest) (*protodef.IncomingTriggerMessageResponse, error)
}

type SessionProcessor interface {
	SetDaryl(*Daryl)

	/* RPC */
	StartWorkSession(*protodef.StartWorkSessionRequest) (*protodef.StartWorkSessionResponse, error)
	CancelWorkSession(*protodef.CancelWorkSessionRequest) (*protodef.CancelWorkSessionResponse, error)
	RefuseSessionSlice(*protodef.RefuseSessionSliceRequest) (*protodef.RefuseSessionSliceResponse, error)
	GetWorkSession(r *protodef.GetWorkSessionRequest) (*protodef.GetWorkSessionResponse, error)
}
