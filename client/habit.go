package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vitaminwater/daryl/protodef"
)

func addHabit(client protodef.DarylServiceClient, c *cli.Context) {
	log.Info("addHabit")

	request := &protodef.AddHabitRequest{
		DarylIdentifier: c.String("identifier"),
		Habit: &protodef.Habit{
			Id:       c.String("id"),
			Title:    c.String("title"),
			Duration: c.String("duration"),
		},
	}
	response, err := client.AddHabit(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func getHabits(client protodef.DarylServiceClient, c *cli.Context) {
	log.Info("getHabits")

	request := &protodef.GetHabitsRequest{
		DarylIdentifier: c.String("identifier"),
	}
	response, err := client.GetHabits(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func habitCommand() cli.Command {
	return cli.Command{
		Name:    "habit",
		Aliases: []string{"h"},
		Subcommands: []cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add a new habit",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "identifier, i",
						Usage: "Daryl's identifier",
					},
					cli.StringFlag{
						Name:  "id",
						Usage: "Habit identifier",
					},
					cli.StringFlag{
						Name:  "title, t",
						Usage: "Title",
					},
					//	Usage: "Cron line\nex: */2 * * * * *\nex: @hourly\nex: @every 1h30m",
					cli.StringFlag{
						Name:  "duration, d",
						Usage: "Duration\nex: 2h45m",
					},
				},
				Action: func(c *cli.Context) error {
					_, daryl := openConnection(c)
					addHabit(daryl, c)
					return nil
				},
			},
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "Get habits",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "identifier, i",
						Usage: "Daryl's identifier",
					},
				},
				Action: func(c *cli.Context) error {
					_, daryl := openConnection(c)
					getHabits(daryl, c)
					return nil
				},
			},
		},
	}
}
