package deliver

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
	"github.com/sfomuseum/go-flags/flagset"
)

const STDIN string = "-"

type RunOptions struct {
	// A registered aaronland/go-mailinglist/v2/database.SubscriptionsDatabase URI.
	SubscriptionsDatabaseURI string
	// A registered aaronland/go-mailinglist/v2/database.DeliveriesDatabase URI.
	DeliveriesDatabaseURI string
	// A registered aaronland/go-mailinglist/v2/database.EventLogsDatabase URI.
	EventLogsDatabaseURI string
	// A registered aaronland/go-mail.Sender URI.
	SenderURI string
	// One or more addresses to deliver the message to.
	To []string
	// The address delivering the message. If empty then the message will be delivered to all subscribers whose status is "enabled".
	From string
	// The subject of the message being delivered.
	Subject string
	// The body of the message being delivered.
	Body string
	// The content-type of the message body.
	ContentType string
	// Optional custom message ID to assign to the message. If empty a unique key will be generated on delivery.
	MessageId string
	// Zero or more URIs referencing files to attach to the message. URIs are dereferenced using the aaronland/gocloud-blob/bucket.ParseURI method.
	Attachments []string
	// Enable verbose (debug) logging.
	Verbose bool
}

func RunOptionsFromFlagSet(fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	opts := &RunOptions{
		SubscriptionsDatabaseURI: subscriptions_database_uri,
		DeliveriesDatabaseURI:    deliveries_database_uri,
		EventLogsDatabaseURI:     eventlogs_database_uri,
		SenderURI:                sender_uri,
		To:                       to,
		From:                     from,
		Subject:                  subject,
		Body:                     body,
		MessageId:                message_id,
		ContentType:              content_type,
		Attachments:              attachments,
		Verbose:                  verbose,
	}

	if opts.Body == STDIN {

		body, err := io.ReadAll(os.Stdin)

		if err != nil {
			return nil, fmt.Errorf("Failed to read message body from STDIN, %w", err)
		}

		opts.Body = string(body)
	}

	if opts.MessageId == "" {

		guid := uuid.New()
		opts.MessageId = guid.String()
	}

	return opts, nil
}
