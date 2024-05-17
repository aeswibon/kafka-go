package main

import (
	"log"
	"os"

	"github.com/aeswibon/kafka/internal/consumer"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <topic>", os.Args[0])
	}
	topic := os.Args[1]

	cons, err := consumer.NewConsumer("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}

	if err := cons.Subscribe(topic); err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}
}
