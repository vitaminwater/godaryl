package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/farm"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const Identifier = "lol3"

func startDaryl(client farm.FarmClient) {
	log.Info("startDaryl")
	request := &farm.StartDarylRequest{Identifier: Identifier}
	response, err := client.StartDaryl(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func hasDaryl(client farm.FarmClient) {
	log.Info("hasDaryl")
	request := &farm.HasDarylRequest{Identifier: Identifier}
	response, err := client.HasDaryl(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func userMessage(client daryl.DarylClient) {
	log.Info("userMessage")
	request := &daryl.UserMessageRequest{Identifier: Identifier, Text: "http://lol.com/pouet It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout."}
	response, err := client.UserMessage(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to stuff: %v", err)
	}
	log.Println(response)
}

func addHabit(client daryl.DarylClient) {
	log.Info("addHabit")
	request := &daryl.AddHabitRequest{
		Identifier: Identifier,
		Habit: &daryl.Habit{
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

func startWorkSession(client daryl.DarylClient) {
	log.Info("startWorkSession")
	request := &daryl.StartWorkSessionRequest{Identifier: Identifier}
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

	farm := farm.NewFarmClient(conn)
	daryl := daryl.NewDarylClient(conn)

	startDaryl(farm)
	hasDaryl(farm)
	userMessage(daryl)
	addHabit(daryl)
	addHabit(daryl)
	startWorkSession(daryl)
}
