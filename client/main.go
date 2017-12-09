package main

import (
	"os"

	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vitaminwater/daryl/protodef"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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

func userMessage(client protodef.DarylServiceClient, c *cli.Context) {
	log.Info("userMessage")
	request := &protodef.UserMessageRequest{
		DarylIdentifier: c.String("identifier"),
		Message: &protodef.Message{
			Text:            c.String("message"),
			HabitIdentifier: c.String("habit"),
			At:              ptypes.TimestampNow(),
			Attrs:           []byte("{}"),
		},
	}
	response, err := client.UserMessage(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

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

func startWorkSession(client protodef.DarylServiceClient, c *cli.Context) {
	log.Info("startWorkSession")
	request := &protodef.StartWorkSessionRequest{DarylIdentifier: c.String("identifier"), Config: &protodef.SessionConfig{Duration: c.String("duration")}}
	response, err := client.StartWorkSession(context.Background(), request)
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
			Aliases: []string{"m"},
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
					Name:  "id",
					Usage: "Habit identifier",
				},
				cli.StringFlag{
					Name:  "title, t",
					Usage: "Title",
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
