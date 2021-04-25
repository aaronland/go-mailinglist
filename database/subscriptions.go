package database

import (
	"context"
	"fmt"
	"github.com/aaronland/go-mailinglist/confirmation"
	"github.com/aaronland/go-mailinglist/delivery"
	"github.com/aaronland/go-mailinglist/eventlog"
	"github.com/aaronland/go-mailinglist/invitation"
	"github.com/aaronland/go-mailinglist/subscription"
	"github.com/aaronland/go-roster"
	"net/url"
	"sort"
	"strings"
)

type ListSubscriptionsFunc func(*subscription.Subscription) error
type ListConfirmationsFunc func(*confirmation.Confirmation) error
type ListEventLogsFunc func(*eventlog.EventLog) error
type ListDeliveriesFunc func(*delivery.Delivery) error
type ListInvitationsFunc func(*invitation.Invitation) error

type SubscriptionsDatabase interface {
	AddSubscription(*subscription.Subscription) error
	RemoveSubscription(*subscription.Subscription) error
	UpdateSubscription(*subscription.Subscription) error
	GetSubscriptionWithAddress(string) (*subscription.Subscription, error)
	ListSubscriptions(context.Context, ListSubscriptionsFunc) error
	ListSubscriptionsWithStatus(context.Context, ListSubscriptionsFunc, ...int) error
}

var subscriptions_databases roster.Roster

// The initialization function signature for implementation of the Subscriptions_database interface.
type SubscriptionsDatabaseInitializeFunc func(context.Context, string) (SubscriptionsDatabase, error)

// Ensure that the internal roster.Roster instance has been created successfully.
func ensureSubscriptions_databaseRoster() error {

	if subscriptions_databases == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		subscriptions_databases = r
	}

	return nil
}

// Register a new URI scheme and Subscriptions_databaseInitializeFunc function for a implementation of the Subscriptions_database interface.
func RegisterSubscriptionsDatabase(ctx context.Context, scheme string, f SubscriptionsDatabaseInitializeFunc) error {

	err := ensureSubscriptions_databaseRoster()

	if err != nil {
		return err
	}

	return subscriptions_databases.Register(ctx, scheme, f)
}

// Return a list of URI schemes for registered implementations of the Subscriptions_database interface.
func SubscriptionsDatabasesSchemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureSubscriptions_databaseRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range subscriptions_databases.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}

// Create a new instance of the Subscriptions_database interface. Subscriptions_database instances are created by
// passing in a context.Context instance and a URI string. The form and substance of
// URI strings are specific to their implementations. For example to create a OAuth1Subscriptions_database
// you would write:
// cl, err := subscriptions_database.NewSubscriptions_database(ctx, "oauth1://?consumer_key={KEY}&consumer_secret={SECRET}")

func NewSubscriptionsDatabase(ctx context.Context, uri string) (SubscriptionsDatabase, error) {

	// To account for things that might be gocloud.dev/runtimevar-encoded
	// in a file using editors that automatically add newlines (thanks, Emacs)

	uri = strings.TrimSpace(uri)

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := subscriptions_databases.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	f := i.(SubscriptionsDatabaseInitializeFunc)
	return f(ctx, uri)
}
