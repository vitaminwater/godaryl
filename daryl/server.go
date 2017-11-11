package daryl

import (
	"context"
	"github.com/cskr/pubsub"
	log "github.com/sirupsen/logrus"
	"sync"
)

type darylServer struct {
	pubsub   *pubsub.PubSub
	registry *sync.Map
}

func (s *darylServer) pub(msg interface{}, topic string) {
	s.pubsub.Pub(topicMessage{topic, msg}, topic)
}

func (s *darylServer) UserMessage(c context.Context, r *UserMessageRequest) (*UserMessageResponse, error) {
	log.Info("UserMessage")
	s.pub(r, USER_MESSAGE_TOPIC)
	return &UserMessageResponse{}, nil
}

func (s *darylServer) AddHabit(c context.Context, r *AddHabitRequest) (*AddHabitResponse, error) {
	log.Info("AddHabit")
	s.pub(r, ADD_HABIT_TOPIC)
	return &AddHabitResponse{}, nil
}

func (s *darylServer) StartWorkSession(c context.Context, r *StartWorkSessionRequest) (*StartWorkSessionResponse, error) {
	log.Info("StartWorkSession")
	s.pub(r, START_WORK_SESSION_TOPIC)
	return &StartWorkSessionResponse{}, nil
}

func NewServer(registry *sync.Map) *darylServer {
	s := &darylServer{pubsub.New(10), registry}
	newDarylLog(s)
	newMessageRouter(s)
	return s
}
