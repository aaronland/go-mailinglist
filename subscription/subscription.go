package subscription

import (
	"errors"
	"net/mail"
	"time"
)

const SUBSCRIPTION_STATUS_PENDING int = 0
const SUBSCRIPTION_STATUS_ENABLED int = 1
const SUBSCRIPTION_STATUS_DISABLED int = 2
const SUBSCRIPTION_STATUS_BLOCKED int = 3

type Subscription struct {
	Address      string `json:"address"`
	Created      int64  `json:"created"`
	Confirmed    int64  `json:"confirmed"`
	LastModified int64  `json:"lastmodified"`
	Status       int    `json:"status"`
}

func NewSubscription(str_addr string) (*Subscription, error) {

	addr, err := mail.ParseAddress(str_addr)

	if err != nil {
		return nil, err
	}

	now := time.Now()
	ts := now.Unix()

	sub := &Subscription{
		Address:      addr.Address,
		Created:      ts,
		LastModified: ts,
		Confirmed:    0,
		Status:       SUBSCRIPTION_STATUS_PENDING,
	}

	return sub, nil
}

func (s *Subscription) Confirm() error {
	now := time.Now()
	s.Confirmed = now.Unix()

	s.Status = SUBSCRIPTION_STATUS_ENABLED

	s.setLastModified()
	return nil
}

func (s *Subscription) Unconfirm() error {
	s.Confirmed = 0
	s.Status = SUBSCRIPTION_STATUS_PENDING

	s.setLastModified()
	return nil
}

func (s *Subscription) Enable() error {

	if s.IsBlocked() {
		return errors.New("Subscriber is blocked and needs to be unblocked first")
	}

	s.Status = SUBSCRIPTION_STATUS_ENABLED

	s.setLastModified()
	return nil
}

func (s *Subscription) Disable() error {
	s.Status = SUBSCRIPTION_STATUS_DISABLED

	s.setLastModified()
	return nil
}

func (s *Subscription) Block() error {
	s.Status = SUBSCRIPTION_STATUS_BLOCKED

	s.setLastModified()
	return nil
}

func (s *Subscription) Unblock() error {
	s.Status = SUBSCRIPTION_STATUS_PENDING

	s.setLastModified()
	return nil
}

func (s *Subscription) setLastModified() {
	now := time.Now()
	s.LastModified = now.Unix()
}

func (s *Subscription) IsBlocked() bool {

	if s.Status == SUBSCRIPTION_STATUS_BLOCKED {
		return true
	}

	return false
}

func (s *Subscription) IsEnabled() bool {

	if s.Status == SUBSCRIPTION_STATUS_ENABLED {
		return true
	}

	return false
}

func (s *Subscription) IsConfirmed() bool {

	if s.Confirmed > 0 {
		return true
	}

	return false
}
