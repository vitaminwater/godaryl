package farm

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	context "golang.org/x/net/context"
	"sync"
)

type farmServer struct {
	registry *sync.Map
}

func (f *farmServer) StartDaryl(c context.Context, r *StartDarylRequest) (*StatusResponse, error) {
	log.Println("StartDaryl")
	if _, ok := f.registry.Load(r.Identifier); ok != false {
		return nil, errors.New(fmt.Sprintf("%s already registered", r.Identifier))
	}
	d := daryl.NewDaryl(r.Identifier)
	f.registry.Store(r.Identifier, d)
	return &StatusResponse{true}, nil
}

func (f *farmServer) HasDaryl(c context.Context, r *HasDarylRequest) (*StatusResponse, error) {
	log.Println("HasDaryl")
	_, ok := f.registry.Load(r.Identifier)
	return &StatusResponse{ok}, nil
}

func NewServer(registry *sync.Map) *farmServer {
	s := &farmServer{registry}
	return s
}
