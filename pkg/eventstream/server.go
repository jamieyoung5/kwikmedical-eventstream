package eventstream

import (
	"context"
	"kwikmedical-eventstream/gen"
	v1 "kwikmedical-eventstream/gen/io/cloudevents/v1"
	"log"
	"sync"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	gen.UnimplementedEventStreamV1Server
	mu          sync.Mutex
	subscribers map[string]*Subscriber
}

type Subscriber struct {
	stream     gen.EventStreamV1_SubscribeToEventsServer
	criteria   *gen.SubscriptionRequest
	cancelFunc context.CancelFunc
}

func NewServer() *Server {
	return &Server{
		subscribers: make(map[string]*Subscriber),
	}
}

// PublishEvent handles incoming events and distributes them to subscribers.
func (s *Server) PublishEvent(ctx context.Context, event *v1.CloudEvent) (*gen.PublishEventResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for id, sub := range s.subscribers {
		if matches(sub.criteria, event) {
			if err := sub.stream.Send(event); err != nil {
				log.Printf("Error sending event to subscriber %s: %v", id, err)
			}
		}
	}

	return &gen.PublishEventResponse{
		Id:      event.Id,
		Success: true,
	}, nil
}

// SubscribeToEvents registers a subscriber and starts sending events.
func (s *Server) SubscribeToEvents(req *gen.SubscriptionRequest, stream gen.EventStreamV1_SubscribeToEventsServer) error {
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

	log.Printf("Subscriber %s registered", consumerID)

	<-ctx.Done()

	s.mu.Lock()
	delete(s.subscribers, consumerID)
	s.mu.Unlock()
	log.Printf("Subscriber %s disconnected", consumerID)
	return nil
}

// Unsubscribe removes a subscriber from the list.
func (s *Server) Unsubscribe(ctx context.Context, req *gen.UnsubscribeRequest) (*emptypb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if sub, exists := s.subscribers[req.ConsumerId]; exists {
		sub.cancelFunc()
		delete(s.subscribers, req.ConsumerId)
		log.Printf("Subscriber %s unsubscribed", req.ConsumerId)
	}
	return &emptypb.Empty{}, nil
}
