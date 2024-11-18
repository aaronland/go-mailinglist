package database

import (
	"context"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/aaronland/go-mailinglist/v2/invitation"
	"github.com/aaronland/go-mailinglist/v2/subscription"
	"github.com/aaronland/go-roster"
)

type ListInvitationsFunc func(*invitation.Invitation) error

type InvitationsDatabase interface {
	AddInvitation(context.Context, *invitation.Invitation) error
	RemoveInvitation(context.Context, *invitation.Invitation) error
	UpdateInvitation(context.Context, *invitation.Invitation) error
	GetInvitationWithCode(context.Context, string) (*invitation.Invitation, error)
	GetInvitationWithInvitee(context.Context, string) (*invitation.Invitation, error)
	ListInvitations(context.Context, ListInvitationsFunc) error
	ListInvitationsWithInviter(context.Context, ListInvitationsFunc, *subscription.Subscription) error
	Close() error
}

var invitations_database_roster roster.Roster

// InvitationsDatabaseInitializationFunc is a function defined by individual invitations_database package and used to create
// an instance of that invitations_database
type InvitationsDatabaseInitializationFunc func(ctx context.Context, uri string) (InvitationsDatabase, error)

// RegisterInvitationsDatabase registers 'scheme' as a key pointing to 'init_func' in an internal lookup table
// used to create new `InvitationsDatabase` instances by the `NewInvitationsDatabase` method.
func RegisterInvitationsDatabase(ctx context.Context, scheme string, init_func InvitationsDatabaseInitializationFunc) error {

	err := ensureInvitationsDatabaseRoster()

	if err != nil {
		return err
	}

	return invitations_database_roster.Register(ctx, scheme, init_func)
}

func ensureInvitationsDatabaseRoster() error {

	if invitations_database_roster == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		invitations_database_roster = r
	}

	return nil
}

// NewInvitationsDatabase returns a new `InvitationsDatabase` instance configured by 'uri'. The value of 'uri' is parsed
// as a `url.URL` and its scheme is used as the key for a corresponding `InvitationsDatabaseInitializationFunc`
// function used to instantiate the new `InvitationsDatabase`. It is assumed that the scheme (and initialization
// function) have been registered by the `RegisterInvitationsDatabase` method.
func NewInvitationsDatabase(ctx context.Context, uri string) (InvitationsDatabase, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := invitations_database_roster.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	init_func := i.(InvitationsDatabaseInitializationFunc)
	return init_func(ctx, uri)
}

// Schemes returns the list of schemes that have been registered.
func InvitationsDatabaseSchemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureInvitationsDatabaseRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range invitations_database_roster.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}
