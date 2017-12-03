package main

import (
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/protodef"
	"google.golang.org/grpc"
)

var daryls = sync.Map{}

func openDarylConnection(url string) (protodef.DarylServiceClient, func()) {
	if d, ok := daryls.Load(url); ok == true {
		c := d.(*sync.Pool).Get().(protodef.DarylServiceClient)
		return c, func() { d.(*sync.Pool).Put(c) }
	}
	d := &sync.Pool{
		New: func() interface{} {
			log.Info("openDarylConnection")
			conn, err := grpc.Dial(url, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("fail to dial: %v", err)
			}
			daryl := protodef.NewDarylServiceClient(conn)
			return daryl
		},
	}
	daryls.Store(url, d)
	c := d.Get().(protodef.DarylServiceClient)
	return c, func() { d.Put(c) }
}

var farms = sync.Map{}

func openFarmConnection(url string) (protodef.FarmServiceClient, func()) {
	if f, ok := farms.Load(url); ok == true {
		c := f.(*sync.Pool).Get().(protodef.FarmServiceClient)
		return c, func() { f.(*sync.Pool).Put(c) }
	}
	f := &sync.Pool{
		New: func() interface{} {
			log.Info("openFarmConnection")
			conn, err := grpc.Dial(url, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("fail to dial: %v", err)
			}
			farm := protodef.NewFarmServiceClient(conn)
			return farm
		},
	}
	farms.Store(url, f)
	c := f.Get().(protodef.FarmServiceClient)
	return c, func() { f.Put(c) }
}
