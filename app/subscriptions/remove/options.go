package remove

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

type RunOptions struct {
	SubscriptionsDatabaseURI string
	Addresses                []string
	Verbose                  bool
}

func RunOptionsFromFlagSet(fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	opts := &RunOptions{
		SubscriptionsDatabaseURI: subscriptions_database_uri,
		Addresses:                addresses,
		Verbose:                  verbose,
	}

	return opts, nil
}