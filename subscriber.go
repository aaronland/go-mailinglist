package mailinglist

import (
	"net/mail"
	"time"
)

type Subscriber struct {
	Address   string `json:"address"`
	Created   int64  `json:"created"`
	Confirmed bool   `json:"confirmed"`
	// Status int `json:"status"`
}

func NewSubscriber(str_addr string) (*Subscriber, error) {

	addr, err := mail.ParseAddress(str_addr)

	if err != nil {
		return nil, err
	}

	now := time.Now()

	sub := &Subscriber{
		Address:   addr.Address,
		Created:   now.Unix(),
		Confirmed: false,
	}

	return sub, nil
}
