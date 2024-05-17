package main

import (
	"log"
	"os"
	"strconv"

	"github.com/aeswibon/kafka/internal/producer"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("Usage: %s <topic> <message> <priority>", os.Args[0])
	}
	topic := os.Args[1]
	message := os.Args[2]
	priority, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalf("Invalid priority: %v", err)
	}

	prod, err := producer.NewProducer("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}

	if err := prod.Publish(topic, message, int32(priority)); err != nil {
		log.Fatalf("Failed to produce message: %v", err)
	}
}
