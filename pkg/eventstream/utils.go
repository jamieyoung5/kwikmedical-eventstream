package eventstream

import (
	"fmt"
	"kwikmedical-eventstream/gen"
	v1 "kwikmedical-eventstream/gen/io/cloudevents/v1"
	"time"
)

// matches checks if the event matches the subscriber's criteria.
func matches(criteria *gen.SubscriptionRequest, event *v1.CloudEvent) bool {
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
