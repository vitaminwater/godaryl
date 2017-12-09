package protodef

import (
	"sync"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var daryls = sync.Map{}

func OpenDarylConnection(url string) (DarylServiceClient, func()) {
	if d, ok := daryls.Load(url); ok == true {
		c := d.(*sync.Pool).Get().(DarylServiceClient)
		return c, func() { d.(*sync.Pool).Put(c) }
	}
	d := &sync.Pool{
		New: func() interface{} {
			log.Info("openDarylConnection")
			conn, err := grpc.Dial(url, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("fail to dial: %v", err)
			}
			daryl := NewDarylServiceClient(conn)
			return daryl
		},
	}
	daryls.Store(url, d)
	c := d.Get().(DarylServiceClient)
	return c, func() { d.Put(c) }
}

var farms = sync.Map{}

func OpenFarmConnection(url string) (FarmServiceClient, func()) {
	if f, ok := farms.Load(url); ok == true {
		c := f.(*sync.Pool).Get().(FarmServiceClient)
		return c, func() { f.(*sync.Pool).Put(c) }
	}
	f := &sync.Pool{
		New: func() interface{} {
			log.Info("openFarmConnection")
			conn, err := grpc.Dial(url, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("fail to dial: %v", err)
			}
			farm := NewFarmServiceClient(conn)
			return farm
		},
	}
	farms.Store(url, f)
	c := f.Get().(FarmServiceClient)
	return c, func() { f.Put(c) }
}
