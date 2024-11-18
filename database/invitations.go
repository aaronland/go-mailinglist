package database

import (
	"context"

	"github.com/aaronland/go-mailinglist/v2/invitation"
	"github.com/aaronland/go-mailinglist/v2/subscription"	
)

type ListInvitationsFunc func(*invitation.Invitation) error

type InvitationsDatabase interface {
	AddInvitation(context.Context, *invitation.Invitation) error
	RemoveInvitation(context.Context, *invitation.Invitation) error
	UpdateInvitation(context.Context, *invitation.Invitation) error
	GetInvitationWithCode(string) (context.Context, *invitation.Invitation, error)
	GetInvitationWithInvitee(string) (context.Context, *invitation.Invitation, error)
	ListInvitations(context.Context, ListInvitationsFunc) error
	ListInvitationsWithInviter(context.Context, ListInvitationsFunc, *subscription.Subscription) error
	Close() error
}

