package daryl

import (
	log "github.com/sirupsen/logrus"
)

type darylLog struct {
	c chan interface{}
}

func darylLogProcess(dl *darylLog) {
	for msg := range dl.c {
		tm := msg.(topicMessage)
		log.Info(tm.topic, " : ", tm.msg)
	}
}

func newDarylLog(s *darylServer) *darylLog {
	dl := &darylLog{}
	dl.c = s.pubsub.Sub(
		USER_MESSAGE_TOPIC,
		ADD_HABIT_TOPIC,
		START_WORK_SESSION_TOPIC,
		LINK_LOG_TOPIC,
		NOTE_LOG_TOPIC,
		TODO_LOG_TOPIC,
		UNPROCESSED_LOG_TOPIC,
	)
	go darylLogProcess(dl)
	return dl
}
