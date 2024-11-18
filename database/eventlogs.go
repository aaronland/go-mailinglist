package database

import (
	"context"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/aaronland/go-mailinglist/v2/eventlog"
	"github.com/aaronland/go-roster"
)

type ListEventLogsFunc func(*eventlog.EventLog) error

type EventLogsDatabase interface {
	AddEventLog(context.Context, *eventlog.EventLog) error
	ListEventLogs(context.Context, ListEventLogsFunc) error
	Close() error
}

var eventlogs_database_roster roster.Roster

// EventLogsDatabaseInitializationFunc is a function defined by individual eventlogs_database package and used to create
// an instance of that eventlogs_database
type EventLogsDatabaseInitializationFunc func(ctx context.Context, uri string) (EventLogsDatabase, error)

// RegisterEventLogsDatabase registers 'scheme' as a key pointing to 'init_func' in an internal lookup table
// used to create new `EventLogsDatabase` instances by the `NewEventLogsDatabase` method.
func RegisterEventLogsDatabase(ctx context.Context, scheme string, init_func EventLogsDatabaseInitializationFunc) error {

	err := ensureEventLogsDatabaseRoster()

	if err != nil {
		return err
	}

	return eventlogs_database_roster.Register(ctx, scheme, init_func)
}

func ensureEventLogsDatabaseRoster() error {

	if eventlogs_database_roster == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		eventlogs_database_roster = r
	}

	return nil
}

// NewEventLogsDatabase returns a new `EventLogsDatabase` instance configured by 'uri'. The value of 'uri' is parsed
// as a `url.URL` and its scheme is used as the key for a corresponding `EventLogsDatabaseInitializationFunc`
// function used to instantiate the new `EventLogsDatabase`. It is assumed that the scheme (and initialization
// function) have been registered by the `RegisterEventLogsDatabase` method.
func NewEventLogsDatabase(ctx context.Context, uri string) (EventLogsDatabase, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := eventlogs_database_roster.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	init_func := i.(EventLogsDatabaseInitializationFunc)
	return init_func(ctx, uri)
}

// Schemes returns the list of schemes that have been registered.
func EventLogsDatabaseSchemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureEventLogsDatabaseRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range eventlogs_database_roster.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}
