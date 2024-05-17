package broker

import (
	"context"
	"sync"

	pb "github.com/aeswibon/kafka/pkg/messaging"
)

// Server represents a broker server
type Server struct {
	pb.UnimplementedBrokerServer
	topics map[string]*priorityQueue
	sync.Mutex
}

// NewServer creates a new server
func NewServer() *Server {
	return &Server{
		topics: make(map[string]*priorityQueue),
	}
}

// Publish lets a client produce messages to a topic
func (s *Server) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	s.Lock()
	if _, exists := s.topics[req.Topic]; !exists {
		s.topics[req.Topic] = &priorityQueue{}
	}
	s.Unlock()
	s.topics[req.Topic].Enqueue(message{content: req.Message, priority: req.Priority})
	return &pb.PublishResponse{Success: true}, nil
}

// Subscribe lets a client consume messages from a topic
func (s *Server) Subscribe(req *pb.SubscribeRequest, stream pb.Broker_SubscribeServer) error {
	if queue, exists := s.topics[req.Topic]; exists {
		for {
			msg := queue.Dequeue()
			if msg == nil {
				break
			}
			if err := stream.Send(&pb.SubscribeResponse{Message: msg.content, Priority: msg.priority}); err != nil {
				return err
			}
		}
	}
	return nil
}
