package consumer

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/aeswibon/kafka/pkg/messaging"
	"google.golang.org/grpc"
)

// Consumer represents a consumer
type Consumer struct {
	client pb.BrokerClient
}

// NewConsumer creates a new consumer
func NewConsumer(addr string) (*Consumer, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	client := pb.NewBrokerClient(conn)
	return &Consumer{client: client}, nil
}

// Subscribe lets a client consume messages from a topic
func (c *Consumer) Subscribe(topic string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	stream, err := c.client.Subscribe(ctx, &pb.SubscribeRequest{Topic: topic})
	if err != nil {
		return err
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("Consumed message from topic %s with priority %d: %s", topic, res.Priority, res.Message)
	}
	return nil
}
