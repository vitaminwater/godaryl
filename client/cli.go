package main

import (
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vitaminwater/daryl/protodef"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func startDaryl(client protodef.FarmClient, c *cli.Context) {
	log.Info("startDaryl")
	request := &protodef.StartDarylRequest{Identifier: c.String("identifier")}
	response, err := client.StartDaryl(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func hasDaryl(client protodef.FarmClient, c *cli.Context) {
	log.Info("hasDaryl")
	request := &protodef.HasDarylRequest{Identifier: c.String("identifier")}
	response, err := client.HasDaryl(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func userMessage(client protodef.DarylClient, c *cli.Context) {
	log.Info("userMessage")
	request := &protodef.UserMessageRequest{
		Identifier: c.String("identifier"),
		Message:    &protodef.Message{Text: c.String("message"), At: ptypes.TimestampNow()}}
	response, err := client.UserMessage(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func addHabit(client protodef.DarylClient, c *cli.Context) {
	log.Info("addHabit")
	deadlinet := time.Time{}
	err := deadlinet.UnmarshalText([]byte(c.String("deadline")))
	if err != nil {
		log.Fatal(err)
	}
	deadline, err := ptypes.TimestampProto(deadlinet)
	if err != nil {
		log.Fatal(err)
	}

	request := &protodef.AddHabitRequest{
		Identifier: c.String("identifier"),
		Habit: &protodef.Habit{
			Title:    c.String("title"),
			Deadline: deadline,
			Cron:     c.String("cron"),
			Duration: c.String("duration"),
			LastDone: ptypes.TimestampNow(),
			Stats:    &protodef.HabitStat{Urgent: 0, NMissed: 0},
		},
	}
	response, err := client.AddHabit(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func startWorkSession(client protodef.DarylClient, c *cli.Context) {
	log.Info("startWorkSession")
	request := &protodef.StartWorkSessionRequest{Identifier: c.String("identifier")}
	response, err := client.StartWorkSession(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func cancelWorkSession(client protodef.DarylClient, c *cli.Context) {
	log.Info("cancelWorkSession")
	request := &protodef.CancelWorkSessionRequest{Identifier: c.String("identifier")}
	response, err := client.CancelWorkSession(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func refuseWorkSession(client protodef.DarylClient, c *cli.Context) {
	log.Info("refuseWorkSession")
	request := &protodef.RefuseSessionSliceRequest{Identifier: c.String("identifier"), Index: &protodef.SessionSliceIndex{uint32(c.Uint("index"))}}
	response, err := client.RefuseSessionSlice(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func openConnection(c *cli.Context) (protodef.FarmClient, protodef.DarylClient) {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	farm := protodef.NewFarmClient(conn)
	daryl := protodef.NewDarylClient(conn)

	return farm, daryl
}

func main() {
	app := cli.NewApp()
	app.Name = "Daryl"
	app.Usage = "Show me what you got"
	app.Commands = []cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "Start a new daryl",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "identifier, i",
					Usage: "Daryl's identifier",
				},
			},
			Action: func(c *cli.Context) error {
				farm, _ := openConnection(c)
				startDaryl(farm, c)
				return nil
			},
		},
		{
			Name:    "hasdaryl",
			Aliases: []string{"hd"},
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
		{
			Name:    "message",
			Aliases: []string{"hd"},
			Usage:   "Sends a message to a daryl",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "identifier, i",
					Usage: "Daryl's identifier",
				},
				cli.StringFlag{
					Name:  "message, m",
					Usage: "Message to send to the Daryl",
				},
			},
			Action: func(c *cli.Context) error {
				_, daryl := openConnection(c)
				userMessage(daryl, c)
				return nil
			},
		},
		{
			Name:    "addhabit",
			Aliases: []string{"a"},
			Usage:   "Add a new habit",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "identifier, i",
					Usage: "Daryl's identifier",
				},
				cli.StringFlag{
					Name:  "title, t",
					Usage: "Title",
				},
				cli.StringFlag{
					Name:  "deadline, d",
					Usage: "Deadline date as RFC 3339.\nex: 2002-10-02T15:00:00Z",
				},
				cli.StringFlag{
					Name:  "cron, c",
					Usage: "Cron line\nex: */2 * * * * *\nex: @hourly\nex: @every 1h30m",
				},
				cli.StringFlag{
					Name:  "duration, l",
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
			Name:    "work",
			Aliases: []string{"w"},
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
					},
					Action: func(c *cli.Context) error {
						_, daryl := openConnection(c)
						startWorkSession(daryl, c)
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
		},
	}

	app.Run(os.Args)
}
