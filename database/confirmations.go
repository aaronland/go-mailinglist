package database

import (
	"context"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/aaronland/go-mailinglist/confirmation"
	"github.com/aaronland/go-roster"
)

type ConfirmationsDatabase interface {
	AddConfirmation(context.Context, *confirmation.Confirmation) error
	RemoveConfirmation(context.Context, *confirmation.Confirmation) error
	GetConfirmationWithCode(context.Context, string) (*confirmation.Confirmation, error)
	ListConfirmations(context.Context, ListConfirmationsFunc) error
}

var confirmations_databases roster.Roster

// The initialization function signature for implementation of the Confirmations_database interface.
type ConfirmationsDatabaseInitializeFunc func(context.Context, string) (ConfirmationsDatabase, error)

// Ensure that the internal roster.Roster instance has been created successfully.
func ensureConfirmations_databaseRoster() error {

	if confirmations_databases == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		confirmations_databases = r
	}

	return nil
}

// Register a new URI scheme and Confirmations_databaseInitializeFunc function for a implementation of the Confirmations_database interface.
func RegisterConfirmationsDatabase(ctx context.Context, scheme string, f ConfirmationsDatabaseInitializeFunc) error {

	err := ensureConfirmations_databaseRoster()

	if err != nil {
		return err
	}

	return confirmations_databases.Register(ctx, scheme, f)
}

// Return a list of URI schemes for registered implementations of the Confirmations_database interface.
func ConfirmationsDatabasesSchemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureConfirmations_databaseRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range confirmations_databases.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}

// Create a new instance of the Confirmations_database interface. Confirmations_database instances are created by
// passing in a context.Context instance and a URI string. The form and substance of
// URI strings are specific to their implementations. For example to create a OAuth1Confirmations_database
// you would write:
// cl, err := confirmations_database.NewConfirmations_database(ctx, "oauth1://?consumer_key={KEY}&consumer_secret={SECRET}")

func NewConfirmationsDatabase(ctx context.Context, uri string) (ConfirmationsDatabase, error) {

	// To account for things that might be gocloud.dev/runtimevar-encoded
	// in a file using editors that automatically add newlines (thanks, Emacs)

	uri = strings.TrimSpace(uri)

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := confirmations_databases.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	f := i.(ConfirmationsDatabaseInitializeFunc)
	return f(ctx, uri)
}
