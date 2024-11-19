package eventstream

import (
	"fmt"
	"github.com/jamieyoung5/kwikmedical-eventstream/pb"
	cloudeventspb "github.com/jamieyoung5/kwikmedical-eventstream/pb/io/cloudevents/v1"
	"time"
)

// matches checks if the event matches the subscriber's criteria.
func matches(criteria *pb.SubscriptionRequest, event *cloudeventspb.CloudEvent) bool {
	if len(criteria.Types) > 0 {
		match := false
		for _, t := range criteria.Types {
			if t == event.Type {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}

	if criteria.Source != "" && criteria.Source != event.Source {
		return false
	}

	return true
}

// generateUniqueConsumerID generates a unique ID for a subscriber.
func generateUniqueConsumerID() string {
	return fmt.Sprintf("consumer-%d", time.Now().UnixNano())
}
