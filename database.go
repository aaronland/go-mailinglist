package mailinglist

import (
	"errors"
)

type SubscriberDatabase interface {
	AddSubscriber(*Subscriber) error
	RemoveSubscriber(*Subscriber) error
	GetSubscribers(func(*Subscriber) error) error
}

func NewSubscriberDatabaseFromDSN(str_dsn string) (SubscriberDatabase, error) {

	return nil, errors.New("Please write me")
}
