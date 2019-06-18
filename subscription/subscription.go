package subscription

import (
	"net/mail"
	"time"
)

type Subscription struct {
	Address   string `json:"address"`
	Created   int64  `json:"created"`
	Confirmed bool   `json:"confirmed"`
	// Status int `json:"status"`
}

func NewSubscription(str_addr string) (*Subscription, error) {

	addr, err := mail.ParseAddress(str_addr)

	if err != nil {
		return nil, err
	}

	now := time.Now()

	sub := &Subscription{
		Address:   addr.Address,
		Created:   now.Unix(),
		Confirmed: false,
	}

	return sub, nil
}
