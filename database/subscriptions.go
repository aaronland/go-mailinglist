package database

import (
	"context"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/aaronland/go-mailinglist/v2/subscription"
	"github.com/aaronland/go-roster"
)

type ListSubscriptionsFunc func(context.Context, *subscription.Subscription) error

type SubscriptionsDatabase interface {
	AddSubscription(context.Context, *subscription.Subscription) error
	RemoveSubscription(context.Context, *subscription.Subscription) error
	UpdateSubscription(context.Context, *subscription.Subscription) error
	GetSubscriptionWithAddress(context.Context, string) (*subscription.Subscription, error)
	ListSubscriptions(context.Context, ListSubscriptionsFunc) error
	ListSubscriptionsWithStatus(context.Context, []int, ListSubscriptionsFunc) error
	Close() error
}

var subscriptions_database_roster roster.Roster

// SubscriptionsDatabaseInitializationFunc is a function defined by individual subscriptions_database package and used to create
// an instance of that subscriptions_database
type SubscriptionsDatabaseInitializationFunc func(ctx context.Context, uri string) (SubscriptionsDatabase, error)

// RegisterSubscriptionsDatabase registers 'scheme' as a key pointing to 'init_func' in an internal lookup table
// used to create new `SubscriptionsDatabase` instances by the `NewSubscriptionsDatabase` method.
func RegisterSubscriptionsDatabase(ctx context.Context, scheme string, init_func SubscriptionsDatabaseInitializationFunc) error {

	err := ensureSubscriptionsDatabaseRoster()

	if err != nil {
		return err
	}

	return subscriptions_database_roster.Register(ctx, scheme, init_func)
}

func ensureSubscriptionsDatabaseRoster() error {

	if subscriptions_database_roster == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		subscriptions_database_roster = r
	}

	return nil
}

// NewSubscriptionsDatabase returns a new `SubscriptionsDatabase` instance configured by 'uri'. The value of 'uri' is parsed
// as a `url.URL` and its scheme is used as the key for a corresponding `SubscriptionsDatabaseInitializationFunc`
// function used to instantiate the new `SubscriptionsDatabase`. It is assumed that the scheme (and initialization
// function) have been registered by the `RegisterSubscriptionsDatabase` method.
func NewSubscriptionsDatabase(ctx context.Context, uri string) (SubscriptionsDatabase, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := subscriptions_database_roster.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	init_func := i.(SubscriptionsDatabaseInitializationFunc)
	return init_func(ctx, uri)
}

// Schemes returns the list of schemes that have been registered.
func SubscriptionsDatabaseSchemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureSubscriptionsDatabaseRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range subscriptions_database_roster.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}
