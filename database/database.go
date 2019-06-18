package database

import (
	"context"
	"errors"
	"github.com/aaronland/go-mailinglist/confirmation"	
	"github.com/aaronland/go-mailinglist/subscription"
	"github.com/aaronland/go-string/dsn"
	"strings"
)

type ListSubscriptionsFunc func(*subscription.Subscription) error
type ListConfirmationsFunc func(*confirmation.Confirmation) error

type SubscriptionDatabase interface {
	AddSubscription(*subscription.Subscription) error
	RemoveSubscription(*subscription.Subscription) error
	UpdateSubscription(*subscription.Subscription) error
	GetSubscriptionWithAddress(string) (*subscription.Subscription, error)
	ConfirmedSubscriptions(context.Context, ListSubscriptionsFunc) error
	UnconfirmedSubscriptions(context.Context, ListSubscriptionsFunc) error
}

type ConfirmationDatabase interface {
	AddConfirmation(*confirmation.Confirmation) error
	RemoveConfirmation(*confirmation.Confirmation) error
	GetConfirmationWithCode(string) (*confirmation.Confirmation, error)
	ListConfirmations(context.Context, ListConfirmationsFunc) error
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
			db, err = NewFSSubscriptionDatabase(root)
		} else {
			err = errors.New("Missing 'root' DSN string")
		}

	default:
		err = errors.New("Invalid database")
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
