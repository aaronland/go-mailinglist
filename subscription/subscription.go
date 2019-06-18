package subscription

import (
	"net/mail"
	"time"
)

type Subscription struct {
	Address   string `json:"address"`
	Created   int64  `json:"created"`
	Confirmed int64  `json:"created"`
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
		Confirmed: 0,
	}

	return sub, nil
}

func (s *Subscription) IsConfirmed() bool {

	if s.Confirmed > 0 {
		return true
	}

	return false
}
