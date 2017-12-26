package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vitaminwater/daryl/protodef"
)

func startWorkSession(client protodef.DarylServiceClient, c *cli.Context) {
	log.Info("startWorkSession")
	request := &protodef.StartWorkSessionRequest{DarylIdentifier: c.String("identifier"), Config: &protodef.SessionConfig{Duration: c.String("duration")}}
	response, err := client.StartWorkSession(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func getWorkSession(client protodef.DarylServiceClient, c *cli.Context) {
	log.Info("getWorkSession")
	request := &protodef.GetWorkSessionRequest{
		DarylIdentifier: c.String("identifier"),
	}
	response, err := client.GetWorkSession(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func cancelWorkSession(client protodef.DarylServiceClient, c *cli.Context) {
	log.Info("cancelWorkSession")
	request := &protodef.CancelWorkSessionRequest{DarylIdentifier: c.String("identifier")}
	response, err := client.CancelWorkSession(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func refuseWorkSession(client protodef.DarylServiceClient, c *cli.Context) {
	log.Info("refuseWorkSession")
	request := &protodef.RefuseSessionSliceRequest{DarylIdentifier: c.String("identifier"), Index: &protodef.SessionSliceIndex{Index: uint32(c.Uint("index"))}}
	response, err := client.RefuseSessionSlice(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func sessionCommand() cli.Command {
	return cli.Command{
		Name:    "session",
		Aliases: []string{"s"},
		Subcommands: []cli.Command{
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "Start a new work session",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "identifier, i",
						Usage: "Daryl's identifier",
					},
					cli.StringFlag{
						Name:  "duration, d",
						Usage: "Work session duration",
					},
				},
				Action: func(c *cli.Context) error {
					_, daryl := openConnection(c)
					startWorkSession(daryl, c)
					return nil
				},
			},
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "Get current work session",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "identifier, i",
						Usage: "Daryl's identifier",
					},
				},
				Action: func(c *cli.Context) error {
					_, daryl := openConnection(c)
					getWorkSession(daryl, c)
					return nil
				},
			},
			{
				Name:    "cancel",
				Aliases: []string{"c"},
				Usage:   "Cancel the current work session",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "identifier, i",
						Usage: "Daryl's identifier",
					},
				},
				Action: func(c *cli.Context) error {
					_, daryl := openConnection(c)
					cancelWorkSession(daryl, c)
					return nil
				},
			},
			{
				Name:    "refuse",
				Aliases: []string{"r"},
				Usage:   "Refuse on habit from the current work session",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "identifier, i",
						Usage: "Daryl's identifier",
					},
					cli.StringFlag{
						Name:  "index, idx",
						Usage: "Habit's index",
					},
				},
				Action: func(c *cli.Context) error {
					_, daryl := openConnection(c)
					refuseWorkSession(daryl, c)
					return nil
				},
			},
		},
	}
}
