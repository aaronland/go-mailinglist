package database

import (
	"context"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/aaronland/go-mailinglist/v2/delivery"
	"github.com/aaronland/go-roster"
)

type ListDeliveriesFunc func(context.Context, *delivery.Delivery) error

type DeliveriesDatabase interface {
	AddDelivery(context.Context, *delivery.Delivery) error
	ListDeliveries(context.Context, ListDeliveriesFunc) error
	GetDeliveryWithAddressAndMessageId(context.Context, string, string) (*delivery.Delivery, error)
	Close() error
}

var deliveries_database_roster roster.Roster

// DeliveriesDatabaseInitializationFunc is a function defined by individual deliveries_database package and used to create
// an instance of that deliveries_database
type DeliveriesDatabaseInitializationFunc func(ctx context.Context, uri string) (DeliveriesDatabase, error)

// RegisterDeliveriesDatabase registers 'scheme' as a key pointing to 'init_func' in an internal lookup table
// used to create new `DeliveriesDatabase` instances by the `NewDeliveriesDatabase` method.
func RegisterDeliveriesDatabase(ctx context.Context, scheme string, init_func DeliveriesDatabaseInitializationFunc) error {

	err := ensureDeliveriesDatabaseRoster()

	if err != nil {
		return err
	}

	return deliveries_database_roster.Register(ctx, scheme, init_func)
}

func ensureDeliveriesDatabaseRoster() error {

	if deliveries_database_roster == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		deliveries_database_roster = r
	}

	return nil
}

// NewDeliveriesDatabase returns a new `DeliveriesDatabase` instance configured by 'uri'. The value of 'uri' is parsed
// as a `url.URL` and its scheme is used as the key for a corresponding `DeliveriesDatabaseInitializationFunc`
// function used to instantiate the new `DeliveriesDatabase`. It is assumed that the scheme (and initialization
// function) have been registered by the `RegisterDeliveriesDatabase` method.
func NewDeliveriesDatabase(ctx context.Context, uri string) (DeliveriesDatabase, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := deliveries_database_roster.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	init_func := i.(DeliveriesDatabaseInitializationFunc)
	return init_func(ctx, uri)
}

// Schemes returns the list of schemes that have been registered.
func DeliveriesDatabaseSchemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureDeliveriesDatabaseRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range deliveries_database_roster.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}
