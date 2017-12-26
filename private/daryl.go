package main

import (
	"context"
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/protodef"
)

type darylServer struct {
	registry *sync.Map
}

func (s *darylServer) UserMessage(c context.Context, r *protodef.UserMessageRequest) (*protodef.UserMessageResponse, error) {
	log.Info("UserMessage")
	d, ok := s.registry.Load(r.DarylIdentifier)
	if ok != true {
		log.Info(r)
		return nil, fmt.Errorf("Unknown Daryl %s", r.DarylIdentifier)
	}
	resp, err := d.(*daryl.Daryl).MessageProcessor.UserMessage(r)
	return resp, err
}

func (s *darylServer) GetUserMessages(c context.Context, r *protodef.GetUserMessagesRequest) (*protodef.GetUserMessagesResponse, error) {
	log.Info("GetUserMessages")
	d, ok := s.registry.Load(r.DarylIdentifier)
	if ok != true {
		log.Info(r)
		return nil, fmt.Errorf("Unknown Daryl %s", r.DarylIdentifier)
	}
	resp, err := d.(*daryl.Daryl).MessageProcessor.GetUserMessages(r)
	return resp, err
}

func (s *darylServer) AddHabit(c context.Context, r *protodef.AddHabitRequest) (*protodef.AddHabitResponse, error) {
	log.Info("AddHabit")
	d, ok := s.registry.Load(r.DarylIdentifier)
	if ok != true {
		return nil, fmt.Errorf("Unknown Daryl %s", r.DarylIdentifier)
	}

	resp, err := d.(*daryl.Daryl).HabitProcessor.AddHabit(r)
	return resp, err
}

func (s *darylServer) GetHabits(c context.Context, r *protodef.GetHabitsRequest) (*protodef.GetHabitsResponse, error) {
	log.Info("GetHabits")
	d, ok := s.registry.Load(r.DarylIdentifier)
	if ok != true {
		return nil, fmt.Errorf("Unknown Daryl %s", r.DarylIdentifier)
	}

	resp, err := d.(*daryl.Daryl).HabitProcessor.GetHabits(r)
	return resp, err
}

func (s *darylServer) AddTrigger(c context.Context, r *protodef.AddTriggerRequest) (*protodef.AddTriggerResponse, error) {
	log.Info("AddTrigger")
	d, ok := s.registry.Load(r.DarylIdentifier)
	if ok != true {
		return nil, fmt.Errorf("Unknown Daryl %s", r.DarylIdentifier)
	}

	resp, err := d.(*daryl.Daryl).TriggerProcessor.AddTrigger(r)
	return resp, err
}

func (s *darylServer) StartWorkSession(c context.Context, r *protodef.StartWorkSessionRequest) (*protodef.StartWorkSessionResponse, error) {
	log.Info("StartWorkSession")
	d, ok := s.registry.Load(r.DarylIdentifier)
	if ok != true {
		return nil, fmt.Errorf("Unknown Daryl %s", r.DarylIdentifier)
	}

	resp, err := d.(*daryl.Daryl).SessionProcessor.StartWorkSession(r)
	return resp, err
}

func (s *darylServer) GetWorkSession(c context.Context, r *protodef.GetWorkSessionRequest) (*protodef.GetWorkSessionResponse, error) {
	log.Info("GetWorkSession")
	d, ok := s.registry.Load(r.DarylIdentifier)
	if ok != true {
		return nil, fmt.Errorf("Unknown Daryl %s", r.DarylIdentifier)
	}

	resp, err := d.(*daryl.Daryl).SessionProcessor.GetWorkSession(r)
	return resp, err
}

func (s *darylServer) CancelWorkSession(ctx context.Context, in *protodef.CancelWorkSessionRequest) (*protodef.CancelWorkSessionResponse, error) {
	log.Info("CancelWorkSession")
	d, ok := s.registry.Load(in.DarylIdentifier)
	if ok != true {
		return nil, fmt.Errorf("Unknown Daryl %s", in.DarylIdentifier)
	}

	resp, err := d.(*daryl.Daryl).SessionProcessor.CancelWorkSession(in)
	return resp, err
}

func (s *darylServer) RefuseSessionSlice(ctx context.Context, in *protodef.RefuseSessionSliceRequest) (*protodef.RefuseSessionSliceResponse, error) {
	log.Info("RefuseSessionSlice")
	d, ok := s.registry.Load(in.DarylIdentifier)
	if ok != true {
		return nil, fmt.Errorf("Unknown Daryl %s", in.DarylIdentifier)
	}

	resp, err := d.(*daryl.Daryl).SessionProcessor.RefuseSessionSlice(in)
	return resp, err
}

func (s *darylServer) Get(c context.Context, r *protodef.GetRequest) (*protodef.GetResponse, error) {
	log.Info("Get")
	d, ok := s.registry.Load(r.DarylIdentifier)
	if ok != true {
		log.Info(r)
		return nil, fmt.Errorf("Unknown Daryl %s", r.DarylIdentifier)
	}
	da, err := d.(*daryl.Daryl).D.ToProtodef()
	if err != nil {
		return nil, err
	}

	return &protodef.GetResponse{Daryl: da}, nil
}

func (s *darylServer) IncomingTriggerMessage(ctx context.Context, in *protodef.IncomingTriggerMessageRequest) (*protodef.IncomingTriggerMessageResponse, error) {
	log.Info("IncomingTriggerMessage")
	d, ok := s.registry.Load(in.DarylIdentifier)
	if ok != true {
		return nil, fmt.Errorf("Unknown Daryl %s", in.DarylIdentifier)
	}

	resp, err := d.(*daryl.Daryl).TriggerProcessor.IncomingMessage(in)
	return resp, err
}

func NewDarylServer(registry *sync.Map) *darylServer {
	s := &darylServer{registry}
	return s
}
