package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vitaminwater/daryl/protodef"
)

func addTrigger(client protodef.DarylServiceClient, c *cli.Context) {
	log.Info("addTrigger")

	request := &protodef.AddTriggerRequest{
		DarylIdentifier: c.String("identifier"),
		Trigger: &protodef.Trigger{
			Id:              c.String("id"),
			HabitIdentifier: c.String("habit"),
			Name:            c.String("name"),
			Engine:          c.String("engine"),
			Params:          []byte(c.String("params")),
		},
	}
	response, err := client.AddTrigger(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func triggerCommand() cli.Command {
	return cli.Command{
		Name:    "trigger",
		Aliases: []string{"t"},
		Subcommands: []cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add a new trigger",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "identifier, i",
						Usage: "Daryl's identifier",
					},
					cli.StringFlag{
						Name:  "habit, ha",
						Usage: "Daryl's habit identifier",
					},
					cli.StringFlag{
						Name:  "id",
						Usage: "Trigger identifier",
					},
					cli.StringFlag{
						Name:  "name, n",
						Usage: "Name",
					},
					cli.StringFlag{
						Name:  "engine, e",
						Usage: "Engine: cron, monitor, webhook",
					},
					cli.StringFlag{
						Name:  "params, p",
						Usage: "Params as json string",
					},
				},
				Action: func(c *cli.Context) error {
					_, daryl := openConnection(c)
					addTrigger(daryl, c)
					return nil
				},
			},
		},
	}
}
