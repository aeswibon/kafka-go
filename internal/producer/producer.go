package producer

import (
	"context"
	"log"
	"time"

	pb "github.com/aeswibon/kafka/pkg/messaging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Producer represents a producer
type Producer struct {
	client pb.BrokerClient
}

// NewProducer creates a new producer
func NewProducer(addr string) (*Producer, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewBrokerClient(conn)
	return &Producer{client: client}, nil
}

// Publish lets a client produce messages to a topic
func (p *Producer) Publish(topic, message string, priority int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := p.client.Publish(ctx, &pb.PublishRequest{Topic: topic, Message: message, Priority: priority})
	if err != nil {
		return err
	}
	log.Printf("Produced message to topic %s with priority %d: %s", topic, priority, message)
	return nil
}
