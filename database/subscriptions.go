package database

import (
	"context"

	"github.com/aaronland/go-mailinglist/v2/subscription"	
)

type ListSubscriptionsFunc func(*subscription.Subscription) error

type SubscriptionsDatabase interface {
	AddSubscription(context.Context, *subscription.Subscription) error
	RemoveSubscription(context.Context, *subscription.Subscription) error
	UpdateSubscription(context.Context, *subscription.Subscription) error
	GetSubscriptionWithAddress(context.Context, string) (*subscription.Subscription, error)
	ListSubscriptions(context.Context, ListSubscriptionsFunc) error
	ListSubscriptionsWithStatus(context.Context, ListSubscriptionsFunc, ...int) error
	Close() error
}
