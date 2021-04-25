package database

import (
	"context"
	"fmt"
	"github.com/aaronland/go-mailinglist/eventlog"
	"github.com/aaronland/go-roster"
	"net/url"
	"sort"
	"strings"
)

type EventLogsDatabase interface {
	AddEventLog(*eventlog.EventLog) error
	ListEventLogs(context.Context, ListEventLogsFunc) error
}

var eventlogs_databases roster.Roster

// The initialization function signature for implementation of the EventLogs_database interface.
type EventLogsDatabaseInitializeFunc func(context.Context, string) (EventLogsDatabase, error)

// Ensure that the internal roster.Roster instance has been created successfully.
func ensureEventLogs_databaseRoster() error {

	if eventlogs_databases == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		eventlogs_databases = r
	}

	return nil
}

// Register a new URI scheme and EventLogs_databaseInitializeFunc function for a implementation of the EventLogs_database interface.
func RegisterEventLogsDatabase(ctx context.Context, scheme string, f EventLogsDatabaseInitializeFunc) error {

	err := ensureEventLogs_databaseRoster()

	if err != nil {
		return err
	}

	return eventlogs_databases.Register(ctx, scheme, f)
}

// Return a list of URI schemes for registered implementations of the EventLogs_database interface.
func EventLogsDatabasesSchemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureEventLogs_databaseRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range eventlogs_databases.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}

// Create a new instance of the EventLogs_database interface. EventLogs_database instances are created by
// passing in a context.Context instance and a URI string. The form and substance of
// URI strings are specific to their implementations. For example to create a OAuth1EventLogs_database
// you would write:
// cl, err := eventlogs_database.NewEventLogs_database(ctx, "oauth1://?consumer_key={KEY}&consumer_secret={SECRET}")

func NewEventLogsDatabase(ctx context.Context, uri string) (EventLogsDatabase, error) {

	// To account for things that might be gocloud.dev/runtimevar-encoded
	// in a file using editors that automatically add newlines (thanks, Emacs)

	uri = strings.TrimSpace(uri)

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := eventlogs_databases.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	f := i.(EventLogsDatabaseInitializeFunc)
	return f(ctx, uri)
}
