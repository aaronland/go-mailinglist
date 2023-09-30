package database

import (
	"context"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/aaronland/go-mailinglist/delivery"
	"github.com/aaronland/go-roster"
)

type DeliveriesDatabase interface {
	AddDelivery(context.Context, *delivery.Delivery) error
	ListDeliveries(context.Context, ListDeliveriesFunc) error
	GetDeliveryWithAddressAndMessageId(context.Context, string, string) (*delivery.Delivery, error)
}

var deliveries_databases roster.Roster

// The initialization function signature for implementation of the Deliveries_database interface.
type DeliveriesDatabaseInitializeFunc func(context.Context, string) (DeliveriesDatabase, error)

// Ensure that the internal roster.Roster instance has been created successfully.
func ensureDeliveries_databaseRoster() error {

	if deliveries_databases == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		deliveries_databases = r
	}

	return nil
}

// Register a new URI scheme and Deliveries_databaseInitializeFunc function for a implementation of the Deliveries_database interface.
func RegisterDeliveriesDatabase(ctx context.Context, scheme string, f DeliveriesDatabaseInitializeFunc) error {

	err := ensureDeliveries_databaseRoster()

	if err != nil {
		return err
	}

	return deliveries_databases.Register(ctx, scheme, f)
}

// Return a list of URI schemes for registered implementations of the Deliveries_database interface.
func DeliveriesDatabasesSchemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureDeliveries_databaseRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range deliveries_databases.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}

// Create a new instance of the Deliveries_database interface. Deliveries_database instances are created by
// passing in a context.Context instance and a URI string. The form and substance of
// URI strings are specific to their implementations. For example to create a OAuth1Deliveries_database
// you would write:
// cl, err := deliveries_database.NewDeliveries_database(ctx, "oauth1://?consumer_key={KEY}&consumer_secret={SECRET}")

func NewDeliveriesDatabase(ctx context.Context, uri string) (DeliveriesDatabase, error) {

	// To account for things that might be gocloud.dev/runtimevar-encoded
	// in a file using editors that automatically add newlines (thanks, Emacs)

	uri = strings.TrimSpace(uri)

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := deliveries_databases.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	f := i.(DeliveriesDatabaseInitializeFunc)
	return f(ctx, uri)
}
