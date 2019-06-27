package eventlog

import (
	"github.com/aaronland/go-mailinglist/subscription"
	"time"
)

type EventLog struct {
	Address string `json:"address"`
	Created int64  `json:"created"`
	Event   string `json:"event"`
	Message string `json:"message"`
}

func NewEventLogWithSubscription(sub *subscription.Subscription, event string, message string) (*EventLog, error) {

	now := time.Now()

	e := &EventLog{
		Address: sub.Address,
		Created: now.Unix(),
		Event:   event,
		Message: message,
	}

	return e, nil
}
