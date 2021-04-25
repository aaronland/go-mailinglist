package database

import (
	"context"
	"github.com/aaronland/go-mailinglist/invitation"
	"github.com/aaronland/go-mailinglist/subscription"
)

type InvitationsDatabase interface {
	AddInvitation(*invitation.Invitation) error
	RemoveInvitation(*invitation.Invitation) error
	UpdateInvitation(*invitation.Invitation) error
	GetInvitationWithCode(string) (*invitation.Invitation, error)
	GetInvitationWithInvitee(string) (*invitation.Invitation, error)
	ListInvitations(context.Context, ListInvitationsFunc) error
	ListInvitationsWithInviter(context.Context, ListInvitationsFunc, *subscription.Subscription) error
}
