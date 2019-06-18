package database

import (
	"context"
	"errors"
	"github.com/aaronland/go-mailinglist/subscription"
)

type SubscriptionDatabase interface {
	AddSubscription(*subscription.Subscription) error
	RemoveSubscription(*subscription.Subscription) error
	UpdateSubscription(*subscription.Subscription) error
	GetSubscriptionWithAddress(string) (*subscription.Subscription, error)
	ConfirmedSubscriptions(context.Context, func(*subscription.Subscription) error) error
	UnconfirmedSubscriptions(context.Context, func(*subscription.Subscription) error) error
}

func NewSubscriptionDatabaseFromDSN(str_dsn string) (SubscriptionDatabase, error) {

	dsn_map, err := dsn.StringToDSNWithKeys(str_dsn, "database")

	if err != nil {
		return nil, err
	}

	var db SubscriptionDatabase

	switch strings.ToUpper(dsn_map["database"]) {
	case "FS":

		root, ok := dsn_map["root"]

		if ok {
			db, err = NewFSSubscriptionDatabase(str_dsn)
		} else {
			err = errors.New("Missing 'root' DSN string")
		}

	default:
		err = errors.New("Invalid sender")
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
