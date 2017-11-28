package main

import (
	"fmt"
	"net"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vitaminwater/daryl/config"
	"github.com/vitaminwater/daryl/db"
	"github.com/vitaminwater/daryl/distributed"
	"github.com/vitaminwater/daryl/kv"
	"github.com/vitaminwater/daryl/protodef"
	"google.golang.org/grpc"
)

func startServer() {
	lis, err := net.Listen("tcp", config.AppContext.String("bind-string"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	registry := &sync.Map{}

	grpcServer := grpc.NewServer()
	protodef.RegisterFarmServiceServer(grpcServer, NewFarmServer(registry))
	protodef.RegisterDarylServiceServer(grpcServer, NewDarylServer(registry))
	log.Infof("Serving on %s", config.AppContext.String("bind-string"))
	go distributed.Beacon(fmt.Sprintf("private_%s", config.AppContext.String("advertized-url")), config.AppContext.String("advertized-url"))
	grpcServer.Serve(lis)
}

func main() {
	app := cli.NewApp()
	app.Name = "Daryl private server"
	app.Usage = "Show me what you got"
	app.Flags = config.Flags
	app.Action = func(c *cli.Context) error {
		config.AppContext = c
		distributed.Init()
		daryl_db.Init()
		kv.Init()
		startServer()
		return nil
	}
	app.Run(os.Args)
}
