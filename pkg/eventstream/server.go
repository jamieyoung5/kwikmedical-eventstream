package eventstream

import (
	"context"
	"github.com/jamieyoung5/kwikmedical-eventstream/pb"
	cloudeventspb "github.com/jamieyoung5/kwikmedical-eventstream/pb/io/cloudevents/v1"
	"go.uber.org/zap"
	"log"
	"sync"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedEventStreamV1Server
	mu          sync.Mutex
	subscribers map[string]*Subscriber
	logger      *zap.Logger
}

type Subscriber struct {
	stream     pb.EventStreamV1_SubscribeToEventsServer
	criteria   *pb.SubscriptionRequest
	cancelFunc context.CancelFunc
}

func NewServer(logger *zap.Logger) *Server {
	return &Server{
		subscribers: make(map[string]*Subscriber),
		logger:      logger,
	}
}

// PublishEvent handles incoming events and distributes them to subscribers
func (s *Server) PublishEvent(ctx context.Context, event *cloudeventspb.CloudEvent) (*pb.PublishEventResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for id, sub := range s.subscribers {
		if matches(sub.criteria, event) {
			if err := sub.stream.Send(event); err != nil {
				log.Printf("Error sending event to subscriber %s: %v", id, err)
			}
		}
	}

	s.logger.Debug("successfully published event", zap.String("event", event.String()))

	return &pb.PublishEventResponse{
		Id:      event.Id,
		Success: true,
	}, nil
}

// SubscribeToEvents registers a subscriber and starts sending events
func (s *Server) SubscribeToEvents(req *pb.SubscriptionRequest, stream pb.EventStreamV1_SubscribeToEventsServer) error {
	s.mu.Lock()
	consumerID := req.ConsumerId
	if consumerID == "" {
		consumerID = generateUniqueConsumerID()
	}

	ctx, cancel := context.WithCancel(stream.Context())
	sub := &Subscriber{
		stream:     stream,
		criteria:   req,
		cancelFunc: cancel,
	}
	s.subscribers[consumerID] = sub
	s.mu.Unlock()

	s.logger.Debug("Subscriber registered", zap.String("consumer id", consumerID))

	<-ctx.Done()

	s.mu.Lock()
	delete(s.subscribers, consumerID)
	s.mu.Unlock()
	s.logger.Debug("Subscriber disconnected", zap.String("consumer id", consumerID))
	return nil
}

// Unsubscribe removes a subscriber from the list.
func (s *Server) Unsubscribe(ctx context.Context, req *pb.UnsubscribeRequest) (*emptypb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if sub, exists := s.subscribers[req.ConsumerId]; exists {
		sub.cancelFunc()
		delete(s.subscribers, req.ConsumerId)
		s.logger.Debug("Subscriber unsubscribed", zap.String("consumer id", req.ConsumerId))
	}
	return &emptypb.Empty{}, nil
}
