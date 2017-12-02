package main

import (
	"errors"
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/db"
	"github.com/vitaminwater/daryl/habit"
	"github.com/vitaminwater/daryl/message"
	"github.com/vitaminwater/daryl/model"
	"github.com/vitaminwater/daryl/protodef"
	"github.com/vitaminwater/daryl/session"
	context "golang.org/x/net/context"
)

type farmServer struct {
	registry *sync.Map
}

func (f *farmServer) StartDaryl(c context.Context, r *protodef.StartDarylRequest) (*protodef.StartDarylResponse, error) {
	da, err := model.NewDarylFromProtodef(r.Daryl)
	if err != nil {
		return nil, err
	}
	if da.Id == "" {
		err := daryl_db.Insert("daryl", &da)
		if err != nil {
			return nil, err
		}
	} else {
		if _, ok := f.registry.Load(r.Daryl.Id); ok != false {
			return nil, errors.New(fmt.Sprintf("%s already registered", r.Daryl.Id))
		}
	}
	d := daryl.NewDaryl(da, message.NewMessageProcessor(), habit.NewHabitProcessor(), session.NewSessionProcessor())
	f.registry.Store(da.Id, d)
	p, err := da.ToProtodef()
	if err != nil {
		return nil, err
	}
	return &protodef.StartDarylResponse{Daryl: p}, nil
}

func (f *farmServer) HasDaryl(c context.Context, r *protodef.HasDarylRequest) (*protodef.HasDarylResponse, error) {
	log.Println("HasDaryl")
	_, ok := f.registry.Load(r.DarylIdentifier)
	return &protodef.HasDarylResponse{ok}, nil
}

func NewFarmServer(registry *sync.Map) *farmServer {
	s := &farmServer{registry}
	return s
}
