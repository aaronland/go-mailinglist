package database

import (
	"context"
	"github.com/aaronland/go-mailinglist/confirmation"
	"github.com/aaronland/go-mailinglist/subscription"
)

type ListSubscriptionsFunc func(*subscription.Subscription) error
type ListConfirmationsFunc func(*confirmation.Confirmation) error

type SubscriptionsDatabase interface {
	AddSubscription(*subscription.Subscription) error
	RemoveSubscription(*subscription.Subscription) error
	UpdateSubscription(*subscription.Subscription) error
	GetSubscriptionWithAddress(string) (*subscription.Subscription, error)
	ListSubscriptionsConfirmed(context.Context, ListSubscriptionsFunc) error
	ListSubscriptionsUnconfirmed(context.Context, ListSubscriptionsFunc) error
}

type ConfirmationsDatabase interface {
	AddConfirmation(*confirmation.Confirmation) error
	RemoveConfirmation(*confirmation.Confirmation) error
	GetConfirmationWithCode(string) (*confirmation.Confirmation, error)
	ListConfirmations(context.Context, ListConfirmationsFunc) error
}
