package main

import (
	"fmt"
	"net"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/protodef"
	"google.golang.org/grpc"
)

func main() {
	port := 8081
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	registry := &sync.Map{}

	grpcServer := grpc.NewServer()
	protodef.RegisterFarmServiceServer(grpcServer, NewFarmServer(registry))
	protodef.RegisterDarylServiceServer(grpcServer, NewDarylServer(registry))
	log.Infof("Serving on port %d", port)
	grpcServer.Serve(lis)
}
