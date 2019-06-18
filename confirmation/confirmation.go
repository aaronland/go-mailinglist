package confirmation

import (
	"errors"
	"github.com/aaronland/go-mailinglist/subscription"
	"github.com/aaronland/go-string/random"
	"time"
)

type Confirmation struct {
	Action  string `json:"type"`
	Created int64  `json:"created"`
	Code    string `json:"code"`
	Address string `json:"address"`
}

func NewConfirmationForSubscription(sub *subscription.Subscription, action string) (*Confirmation, error) {

	if sub.Confirmed {
		return nil, errors.New("Already confirmed")
	}

	switch action {
	case "subscribe", "unsubscribe":
		// okay
	default:
		return nil, errors.New("Invalid action")
	}

	opts := random.DefaultOptions()
	opts.AlphaNumeric = true
	opts.Chars = 64

	code, err := random.String(opts)

	if err != nil {
		return nil, err
	}

	now := time.Now()

	c := &Confirmation{
		Action:  action,
		Created: now.Unix(),
		Code:    code,
		Address: sub.Address,
	}

	return c, nil
}
