package main

import (
	"errors"
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/habit"
	"github.com/vitaminwater/daryl/message"
	"github.com/vitaminwater/daryl/protodef"
	"github.com/vitaminwater/daryl/session"
	context "golang.org/x/net/context"
)

type farmServer struct {
	registry *sync.Map
}

func (f *farmServer) StartDaryl(c context.Context, r *protodef.StartDarylRequest) (*protodef.StartDarylResponse, error) {
	log.Println("StartDaryl")
	if _, ok := f.registry.Load(r.DarylIdentifier); ok != false {
		return nil, errors.New(fmt.Sprintf("%s already registered", r.DarylIdentifier))
	}
	d := daryl.NewDaryl(r.DarylIdentifier, message.NewMessageProcessor(), habit.NewHabitProcessor(), session.NewSessionProcessor())
	f.registry.Store(r.DarylIdentifier, d)
	return &protodef.StartDarylResponse{}, nil
}

func (f *farmServer) HasDaryl(c context.Context, r *protodef.HasDarylRequest) (*protodef.HasDarylResponse, error) {
	log.Println("HasDaryl")
	_, ok := f.registry.Load(r.DarylIdentifier)
	return &protodef.HasDarylResponse{Response: ok}, nil
}

func NewFarmServer(registry *sync.Map) *farmServer {
	s := &farmServer{registry}
	return s
}
