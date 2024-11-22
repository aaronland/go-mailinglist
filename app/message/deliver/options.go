package deliver

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

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
		Verbose:                  verbose,
	}

	return opts, nil
}
