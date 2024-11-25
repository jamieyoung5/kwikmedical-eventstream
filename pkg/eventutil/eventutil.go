package eventutil

import (
	"context"
	"fmt"
	kwikmedicalpb "github.com/jamieyoung5/kwikmedical-eventstream/pb"
	v1 "github.com/jamieyoung5/kwikmedical-eventstream/pb/io/cloudevents/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func CreateConnection(address string) (kwikmedicalpb.EventStreamV1Client, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
		return nil, err
	}

	client := kwikmedicalpb.NewEventStreamV1Client(conn)
	return client, nil
}

func Publish(
	client kwikmedicalpb.EventStreamV1Client,
	event *v1.CloudEvent,
	logger *zap.Logger,
) (*kwikmedicalpb.PublishEventResponse, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.PublishEvent(ctx, event)
	if err != nil {
		logger.Error("Error publishing event", zap.Error(err))
		return nil, err
	}

	logger.Debug("Event published", zap.String("id", res.Id))
	return res, nil
}

func Consume(
	client kwikmedicalpb.EventStreamV1Client,
	logger *zap.Logger,
	subRequest *kwikmedicalpb.SubscriptionRequest,
	handleEvent func(event *v1.CloudEvent),
) error {
	stream, subErr := client.SubscribeToEvents(context.Background(), subRequest)
	if subErr != nil {
		return fmt.Errorf("failed to subscribe to events: %v", subErr)
	}

	for {
		event, err := stream.Recv()
		if err == io.EOF {
			logger.Debug("EOF received, event stream closed by server.")
			break
		}
		if err != nil {
			return fmt.Errorf("failed to receive events: %v", err)
		}

		handleEvent(event)
	}

	return nil
}
