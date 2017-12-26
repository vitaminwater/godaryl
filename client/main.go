package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vitaminwater/daryl/protodef"
	"google.golang.org/grpc"
)

func openConnection(c *cli.Context) (protodef.FarmServiceClient, protodef.DarylServiceClient) {
	conn, err := grpc.Dial("localhost:8043", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	farm := protodef.NewFarmServiceClient(conn)
	daryl := protodef.NewDarylServiceClient(conn)

	return farm, daryl
}

func main() {
	app := cli.NewApp()
	app.Name = "Daryl"
	app.Usage = "Show me what you got"
	app.Commands = []cli.Command{
		darylCommand(),
		messageCommand(),
		habitCommand(),
		triggerCommand(),
		sessionCommand(),
	}

	app.Run(os.Args)
}
