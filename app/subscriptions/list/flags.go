package list

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var subscriptions_database_uri string
var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("list")

	fs.StringVar(&subscriptions_database_uri, "subscriptions-database-uri", "", "A registered aaronland/go-mailinglist/v2/database.SubscriptionsDatabase URI.")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	return fs
}
