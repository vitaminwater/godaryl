package main

import (
	"fmt"
	"net"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/protodef"
	"github.com/vitaminwater/daryl/server"
	"google.golang.org/grpc"
)

func initEtcd() {
}

func main() {
	port := 8080
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go initEtcd()

	registry := &sync.Map{}

	grpcServer := grpc.NewServer()
	protodef.RegisterFarmServer(grpcServer, server.NewFarmServer(registry))
	protodef.RegisterDarylServer(grpcServer, server.NewDarylServer(registry))
	log.Infof("Serving on port %d", port)
	grpcServer.Serve(lis)
}
