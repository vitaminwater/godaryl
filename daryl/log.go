package daryl

import (
	log "github.com/sirupsen/logrus"
)

type darylLog struct {
	c chan interface{}
}

func darylLogProcess(dl *darylLog) {
	for msg := range dl.c {
		log.Info("################")
		log.Info(msg)
	}
}

func newDarylLog(s *darylServer) *darylLog {
	dl := &darylLog{}
	dl.c = s.pubsub.Sub(USER_MESSAGE_TOPIC)
	go darylLogProcess(dl)
	return dl
}
