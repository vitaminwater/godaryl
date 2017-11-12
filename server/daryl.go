package server

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/protodef"
	"sync"
)

type darylServer struct {
	registry *sync.Map
}

func (s *darylServer) UserMessage(c context.Context, r *protodef.UserMessageRequest) (*protodef.UserMessageResponse, error) {
	log.Info("UserMessage")
	d, ok := s.registry.Load(r.Identifier)
	if ok != true {
		return nil, fmt.Errorf("Unknown Daryl %s", r.Identifier)
	}
	resp, err := d.(*daryl.Daryl).MessageProcessor.UserMessage(r)
	return resp, err
}

func (s *darylServer) AddHabit(c context.Context, r *protodef.AddHabitRequest) (*protodef.AddHabitResponse, error) {
	log.Info("AddHabit")
	d, ok := s.registry.Load(r.Identifier)
	if ok != true {
		return nil, fmt.Errorf("Unknown Daryl %s", r.Identifier)
	}

	resp, err := d.(*daryl.Daryl).HabitProcessor.AddHabit(r)
	return resp, err
}

func (s *darylServer) StartWorkSession(c context.Context, r *protodef.StartWorkSessionRequest) (*protodef.StartWorkSessionResponse, error) {
	log.Info("StartWorkSession")
	d, ok := s.registry.Load(r.Identifier)
	if ok != true {
		return nil, fmt.Errorf("Unknown Daryl %s", r.Identifier)
	}

	resp, err := d.(*daryl.Daryl).SessionProcessor.StartWorkSession(r)
	return resp, err
}

func NewDarylServer(registry *sync.Map) *darylServer {
	s := &darylServer{registry}
	return s
}
