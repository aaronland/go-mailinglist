package database

import (
	"context"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/aaronland/go-mailinglist/invitation"
	"github.com/aaronland/go-mailinglist/subscription"
	"github.com/aaronland/go-roster"
)

type InvitationsDatabase interface {
	AddInvitation(context.Context, *invitation.Invitation) error
	RemoveInvitation(context.Context, *invitation.Invitation) error
	UpdateInvitation(context.Context, *invitation.Invitation) error
	GetInvitationWithCode(context.Context, string) (*invitation.Invitation, error)
	GetInvitationWithInvitee(context.Context, string) (*invitation.Invitation, error)
	ListInvitations(context.Context, ListInvitationsFunc) error
	ListInvitationsWithInviter(context.Context, ListInvitationsFunc, *subscription.Subscription) error
}

var invitations_databases roster.Roster

// The initialization function signature for implementation of the Invitations_database interface.
type InvitationsDatabaseInitializeFunc func(context.Context, string) (InvitationsDatabase, error)

// Ensure that the internal roster.Roster instance has been created successfully.
func ensureInvitations_databaseRoster() error {

	if invitations_databases == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		invitations_databases = r
	}

	return nil
}

// Register a new URI scheme and Invitations_databaseInitializeFunc function for a implementation of the Invitations_database interface.
func RegisterInvitationsDatabase(ctx context.Context, scheme string, f InvitationsDatabaseInitializeFunc) error {

	err := ensureInvitations_databaseRoster()

	if err != nil {
		return err
	}

	return invitations_databases.Register(ctx, scheme, f)
}

// Return a list of URI schemes for registered implementations of the Invitations_database interface.
func InvitationsDatabasesSchemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureInvitations_databaseRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range invitations_databases.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}

// Create a new instance of the Invitations_database interface. Invitations_database instances are created by
// passing in a context.Context instance and a URI string. The form and substance of
// URI strings are specific to their implementations. For example to create a OAuth1Invitations_database
// you would write:
// cl, err := invitations_database.NewInvitations_database(ctx, "oauth1://?consumer_key={KEY}&consumer_secret={SECRET}")

func NewInvitationsDatabase(ctx context.Context, uri string) (InvitationsDatabase, error) {

	// To account for things that might be gocloud.dev/runtimevar-encoded
	// in a file using editors that automatically add newlines (thanks, Emacs)

	uri = strings.TrimSpace(uri)

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := invitations_databases.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	f := i.(InvitationsDatabaseInitializeFunc)
	return f(ctx, uri)
}
