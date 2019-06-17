package mailinglist

import (
	"errors"
)

type SubscriptionDatabase interface {
	AddSubscription(*Subscriber) error
	RemoveSubscription(*Subscriber) error
	Subscriptions(func(*Subscriber) error) error
}

func NewSubscriptionDatabaseFromDSN(str_dsn string) (SubscriptionDatabase, error) {

	return nil, errors.New("Please write me")
}
