package eventlog

import (
	"time"

	"github.com/aaronland/go-mailinglist/v2/subscription"
)

const EVENTLOG_CUSTOM_EVENT int = 0
const EVENTLOG_SUBSCRIBE_EVENT int = 1
const EVENTLOG_UNSUBSCRIBE_EVENT int = 2
const EVENTLOG_ENABLE_EVENT int = 3
const EVENTLOG_DISABLE_EVENT int = 4
const EVENTLOG_BLOCK_EVENT int = 5
const EVENTLOG_UNBLOCK_EVENT int = 6
const EVENTLOG_SEND_OK_EVENT int = 7
const EVENTLOG_SEND_FAIL_EVENT int = 8
const EVENTLOG_CONFIRM_EVENT int = 9
const EVENTLOG_INVITE_REQUEST_EVENT = 10
const EVENTLOG_INVITE_ACCEPT_EVENT = 11

type EventLog struct {
	Address string `json:"address"`
	Created int64  `json:"created"`
	Event   int    `json:"event"`
	Message string `json:"message"`
}

func NewEventLogWithSubscription(sub *subscription.Subscription, event int, message string) (*EventLog, error) {

	now := time.Now()

	e := &EventLog{
		Address: sub.Address,
		Created: now.UnixNano(),
		Event:   event,
		Message: message,
	}

	return e, nil
}
