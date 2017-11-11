package daryl

type topicMessage struct {
	topic string
	msg   interface{}
}

const (
	// From RPC
	USER_MESSAGE_TOPIC       = "USER_MESSAGE_TOPIC"
	ADD_HABIT_TOPIC          = "ADD_HABIT_TOPIC"
	START_WORK_SESSION_TOPIC = "START_WORK_SESSION_TOPIC"

	// Internals
	LINK_LOG_TOPIC        = "LINK_LOG_TOPIC"
	NOTE_LOG_TOPIC        = "NOTE_LOG_TOPIC"
	TODO_LOG_TOPIC        = "TODO_LOG_TOPIC"
	UNPROCESSED_LOG_TOPIC = "UNPROCESSED_LOG_TOPIC"
)
