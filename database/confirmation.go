package database

import (
	"context"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/aaronland/go-mailinglist/v2/confirmation"
	"github.com/aaronland/go-roster"
)

type ListConfirmationsFunc func(*confirmation.Confirmation) error

type ConfirmationsDatabase interface {
	AddConfirmation(context.Context, *confirmation.Confirmation) error
	RemoveConfirmation(context.Context, *confirmation.Confirmation) error
	GetConfirmationWithCode(context.Context, string) (*confirmation.Confirmation, error)
	ListConfirmations(context.Context, ListConfirmationsFunc) error
	Close() error
}

var confirmations_database_roster roster.Roster

// ConfirmationsDatabaseInitializationFunc is a function defined by individual confirmations_database package and used to create
// an instance of that confirmations_database
type ConfirmationsDatabaseInitializationFunc func(ctx context.Context, uri string) (ConfirmationsDatabase, error)

// RegisterConfirmationsDatabase registers 'scheme' as a key pointing to 'init_func' in an internal lookup table
// used to create new `ConfirmationsDatabase` instances by the `NewConfirmationsDatabase` method.
func RegisterConfirmationsDatabase(ctx context.Context, scheme string, init_func ConfirmationsDatabaseInitializationFunc) error {

	err := ensureConfirmationsDatabaseRoster()

	if err != nil {
		return err
	}

	return confirmations_database_roster.Register(ctx, scheme, init_func)
}

func ensureConfirmationsDatabaseRoster() error {

	if confirmations_database_roster == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		confirmations_database_roster = r
	}

	return nil
}

// NewConfirmationsDatabase returns a new `ConfirmationsDatabase` instance configured by 'uri'. The value of 'uri' is parsed
// as a `url.URL` and its scheme is used as the key for a corresponding `ConfirmationsDatabaseInitializationFunc`
// function used to instantiate the new `ConfirmationsDatabase`. It is assumed that the scheme (and initialization
// function) have been registered by the `RegisterConfirmationsDatabase` method.
func NewConfirmationsDatabase(ctx context.Context, uri string) (ConfirmationsDatabase, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := confirmations_database_roster.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	init_func := i.(ConfirmationsDatabaseInitializationFunc)
	return init_func(ctx, uri)
}

// Schemes returns the list of schemes that have been registered.
func ConfirmationsDatabaseSchemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureConfirmationsDatabaseRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range confirmations_database_roster.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}
