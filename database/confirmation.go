package database

import (
	"context"

	"github.com/aaronland/go-mailinglist/v2/confirmation"	
)

type ListConfirmationsFunc func(*confirmation.Confirmation) error

type ConfirmationsDatabase interface {
	AddConfirmation(context.Context, *confirmation.Confirmation) error
	RemoveConfirmation(context.Context, *confirmation.Confirmation) error
	GetConfirmationWithCode(context.Context, string) (*confirmation.Confirmation, error)
	ListConfirmations(context.Context, ListConfirmationsFunc) error
	Close() error
}
