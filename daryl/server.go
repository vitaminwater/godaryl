package daryl

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
)

type darylServer struct {
	registry *sync.Map
}

func (s *darylServer) UserMessage(c context.Context, r *UserMessageRequest) (*UserMessageResponse, error) {
	log.Info("UserMessage")
	d, ok := s.registry.Load(r.Identifier)
	if ok != true {
		return nil, fmt.Errorf("Unknown Daryl %s", r.Identifier)
	}
	d.(*Daryl).Pub(r, USER_MESSAGE_TOPIC)
	return &UserMessageResponse{}, nil
}

func (s *darylServer) AddHabit(c context.Context, r *AddHabitRequest) (*AddHabitResponse, error) {
	log.Info("AddHabit")
	d, ok := s.registry.Load(r.Identifier)
	if ok != true {
		return nil, fmt.Errorf("Unknown Daryl %s", r.Identifier)
	}

	d.(*Daryl).Pub(r, ADD_HABIT_TOPIC)
	return &AddHabitResponse{}, nil
}

func (s *darylServer) StartWorkSession(c context.Context, r *StartWorkSessionRequest) (*StartWorkSessionResponse, error) {
	log.Info("StartWorkSession")
	d, ok := s.registry.Load(r.Identifier)
	if ok != true {
		return nil, fmt.Errorf("Unknown Daryl %s", r.Identifier)
	}

	d.(*Daryl).Pub(r, START_WORK_SESSION_TOPIC)
	return &StartWorkSessionResponse{}, nil
}

func NewServer(registry *sync.Map) *darylServer {
	s := &darylServer{registry}
	return s
}
