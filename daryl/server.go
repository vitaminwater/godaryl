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
	resp, err := d.(*Daryl).MessageProcessor.UserMessage(r)
	return resp, err
}

func (s *darylServer) AddHabit(c context.Context, r *AddHabitRequest) (*AddHabitResponse, error) {
	log.Info("AddHabit")
	d, ok := s.registry.Load(r.Identifier)
	if ok != true {
		return nil, fmt.Errorf("Unknown Daryl %s", r.Identifier)
	}

	resp, err := d.(*Daryl).HabitProcessor.AddHabit(r)
	return resp, err
}

func (s *darylServer) StartWorkSession(c context.Context, r *StartWorkSessionRequest) (*StartWorkSessionResponse, error) {
	log.Info("StartWorkSession")
	d, ok := s.registry.Load(r.Identifier)
	if ok != true {
		return nil, fmt.Errorf("Unknown Daryl %s", r.Identifier)
	}

	resp, err := d.(*Daryl).SessionProcessor.StartWorkSession(r)
	return resp, err
}

func NewServer(registry *sync.Map) *darylServer {
	s := &darylServer{registry}
	return s
}
