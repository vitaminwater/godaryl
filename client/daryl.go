package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vitaminwater/daryl/protodef"
)

func startDaryl(client protodef.FarmServiceClient, c *cli.Context) {
	log.Info("startDaryl")
	request := &protodef.StartDarylRequest{
		Daryl: &protodef.Daryl{
			Id:       c.String("identifier"),
			Name:     c.String("name"),
			Password: c.String("password"),
		},
	}
	response, err := client.StartDaryl(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func hasDaryl(client protodef.FarmServiceClient, c *cli.Context) {
	log.Info("hasDaryl")
	request := &protodef.HasDarylRequest{DarylIdentifier: c.String("identifier")}
	response, err := client.HasDaryl(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func darylCommand() cli.Command {
	return cli.Command{
		Name:    "daryl",
		Aliases: []string{"d"},
		Subcommands: []cli.Command{
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "Start a new daryl",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "identifier, i",
						Usage: "Daryl's identifier",
					},
					cli.StringFlag{
						Name:  "name, n",
						Value: "Daryl",
						Usage: "Daryl's name",
					},
					cli.StringFlag{
						Name:  "password, p",
						Usage: "Daryl's password",
					},
				},
				Action: func(c *cli.Context) error {
					farm, _ := openConnection(c)
					startDaryl(farm, c)
					return nil
				},
			},
			{
				Name:    "exists",
				Aliases: []string{"e"},
				Usage:   "Check if daryl exists here",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "identifier, i",
						Usage: "Daryl's identifier",
					},
				},
				Action: func(c *cli.Context) error {
					farm, _ := openConnection(c)
					hasDaryl(farm, c)
					return nil
				},
			},
		},
	}
}
