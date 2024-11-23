package deliver

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/sfomuseum/go-flags/flagset"
)

const STDIN string = "-"

type RunOptions struct {
	SubscriptionsDatabaseURI string
	DeliveriesDatabaseURI    string
	EventLogsDatabaseURI     string
	SenderURI                string
	To                       []string
	From                     string
	Subject                  string
	Body                     string
	Verbose                  bool
	ContentType              string
	Attachments              []string
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

	return opts, nil
}
