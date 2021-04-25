package sender

import (
	"context"
	"fmt"
	"github.com/aaronland/go-roster"
	"github.com/aaronland/gomail/v2"
	"net/url"
	"sort"
	"strings"
)

var senders roster.Roster

// The initialization function signature for implementation of the Sender interface.
type SenderInitializeFunc func(context.Context, string) (gomail.Sender, error)

// Ensure that the internal roster.Roster instance has been created successfully.
func ensureSenderRoster() error {

	if senders == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		senders = r
	}

	return nil
}

// Register a new URI scheme and SenderInitializeFunc function for a implementation of the Sender interface.
func RegisterSender(ctx context.Context, scheme string, f SenderInitializeFunc) error {

	err := ensureSenderRoster()

	if err != nil {
		return err
	}

	return senders.Register(ctx, scheme, f)
}

// Return a list of URI schemes for registered implementations of the Sender interface.
func Schemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureSenderRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range senders.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}

// Create a new instance of the Sender interface. Sender instances are created by
// passing in a context.Context instance and a URI string. The form and substance of
// URI strings are specific to their implementations.
func NewSender(ctx context.Context, uri string) (gomail.Sender, error) {

	// To account for things that might be gocloud.dev/runtimevar-encoded
	// in a file using editors that automatically add newlines (thanks, Emacs)

	uri = strings.TrimSpace(uri)

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := senders.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	f := i.(SenderInitializeFunc)
	return f(ctx, uri)
}
