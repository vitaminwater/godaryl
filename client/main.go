package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/protodef"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const Identifier = "lol3"

func startDaryl(client protodef.FarmClient) {
	log.Info("startDaryl")
	request := &protodef.StartDarylRequest{Identifier: Identifier}
	response, err := client.StartDaryl(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func hasDaryl(client protodef.FarmClient) {
	log.Info("hasDaryl")
	request := &protodef.HasDarylRequest{Identifier: Identifier}
	response, err := client.HasDaryl(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func userMessage(client protodef.DarylClient) {
	log.Info("userMessage")
	request := &protodef.UserMessageRequest{Identifier: Identifier, Text: "http://lol.com/pouet It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout."}
	response, err := client.UserMessage(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func addHabit(client protodef.DarylClient) {
	log.Info("addHabit")
	request := &protodef.AddHabitRequest{
		Identifier: Identifier,
		Habit: &protodef.Habit{
			Title:       "Habit de ouf !",
			AvgDuration: 30,
			Deadline:    "2002-10-02T15:00:00Z",
			During:      30,
			Every:       2,
			EveryUnit:   "hours",
		},
	}
	response, err := client.AddHabit(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func startWorkSession(client protodef.DarylClient) {
	log.Info("startWorkSession")
	request := &protodef.StartWorkSessionRequest{Identifier: Identifier}
	response, err := client.StartWorkSession(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	farm := protodef.NewFarmClient(conn)
	daryl := protodef.NewDarylClient(conn)

	startDaryl(farm)
	hasDaryl(farm)
	userMessage(daryl)
	addHabit(daryl)
	addHabit(daryl)
	startWorkSession(daryl)
}
