package main

import (
	"log"

	"github.com/vitaminwater/daryl/protodef"
	"google.golang.org/grpc"
)

var daryls map[string]protodef.DarylServiceClient = make(map[string]protodef.DarylServiceClient)

func openDarylConnection(url string) protodef.DarylServiceClient {
	if d, ok := daryls[url]; ok == true {
		return d
	}
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	daryl := protodef.NewDarylServiceClient(conn)
	daryls[url] = daryl
	return daryl
}

var farms map[string]protodef.FarmServiceClient = make(map[string]protodef.FarmServiceClient)

func openFarmConnection(url string) protodef.FarmServiceClient {
	if d, ok := farms[url]; ok == true {
		return d
	}
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	farm := protodef.NewFarmServiceClient(conn)
	farms[url] = farm
	return farm
}
