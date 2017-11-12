package daryl

type TopicMessage struct {
	Topic string
	Msg   interface{}
}

const (
	ALL_TOPIC = "ALL"
	// From RPC
	USER_MESSAGE_TOPIC       = "USER_MESSAGE"
	ADD_HABIT_TOPIC          = "ADD_HABIT"
	START_WORK_SESSION_TOPIC = "START_WORK_SESSION"
)
