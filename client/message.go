package main

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vitaminwater/daryl/protodef"
)

func userMessage(client protodef.DarylServiceClient, c *cli.Context) {
	log.Info("userMessage")
	request := &protodef.UserMessageRequest{
		DarylIdentifier: c.String("identifier"),
		Message: &protodef.Message{
			Text:            c.String("message"),
			HabitIdentifier: c.String("habit"),
			At:              ptypes.TimestampNow(),
			Attrs:           []byte(c.String("attrs")),
		},
	}
	response, err := client.UserMessage(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func getUserMessages(client protodef.DarylServiceClient, c *cli.Context) {
	log.Info("getUserMessage")
	request := &protodef.GetUserMessagesRequest{
		DarylIdentifier: c.String("identifier"),
		HabitIdentifier: c.String("habit"),
		Pagination: &protodef.Pagination{
			From: int32(c.Int("from")),
			To:   int32(c.Int("to")),
		},
	}
	response, err := client.GetUserMessages(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func messageCommand() cli.Command {
	return cli.Command{
		Name:    "message",
		Aliases: []string{"m"},
		Subcommands: []cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Sends a message to a daryl",
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
						Name:  "message, m",
						Usage: "Message to send to the Daryl",
					},
					cli.StringFlag{
						Name:  "attrs, a",
						Usage: "Attrs as json string",
					},
				},
				Action: func(c *cli.Context) error {
					_, daryl := openConnection(c)
					userMessage(daryl, c)
					return nil
				},
			},
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "Gets the messages for a Daryl thread",
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
						Name:  "from, f",
						Usage: "From pagination",
					},
					cli.StringFlag{
						Name:  "to, t",
						Usage: "To pagination",
					},
				},
				Action: func(c *cli.Context) error {
					_, daryl := openConnection(c)
					getUserMessages(daryl, c)
					return nil
				},
			},
		},
	}
}
