package main

import (
	"log"

	"github.com/vitaminwater/daryl/protodef"
	"google.golang.org/grpc"
)

var daryls map[string]protodef.DarylClient = make(map[string]protodef.DarylClient)

func openDarylConnection(url string) protodef.DarylClient {
	if d, ok := daryls[url]; ok == true {
		return d
	}
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	daryl := protodef.NewDarylClient(conn)
	daryls[url] = daryl
	return daryl
}
