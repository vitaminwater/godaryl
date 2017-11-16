package main

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/protodef"
	"github.com/vitaminwater/daryl/server"
	"google.golang.org/grpc"
)

func initEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	for {
		// minimum lease TTL is 5-second
		gr, err := cli.Grant(context.TODO(), 2)
		if err != nil {
			log.Fatal(err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_, err = cli.Put(ctx, "sample_key", "sample_value", clientv3.WithLease(gr.ID))
		cancel()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
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
