package app

import (
	"log"
	"os"

	"github.com/nats-io/stan.go"
)

func Publish() {
	sc, err := stan.Connect("microservice", "microservice_b", stan.NatsURL("nats://nats-streaming:4222"))
	if err != nil {
		log.Fatalf("error connection to nats-streaming: %v", err)
	}
	defer sc.Close()

	filePath := "./data.json"
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error read data.json: %v", err)
	}

	if err := sc.Publish("test", jsonData); err != nil {
		log.Printf("error to publish message: %v", err)

		return
	}

	log.Println("message publish successfully!")
}
