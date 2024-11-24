package status

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var subscriptions_database_uri string
var verbose bool

var addresses multi.MultiString
var status string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("list")

	fs.StringVar(&subscriptions_database_uri, "subscriptions-database-uri", "", "A registered aaronland/go-mailinglist/v2/database.SubscriptionsDatabase URI.")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Var(&addresses, "address", "One or more addresses whose subscriptions should be removed.")
	fs.StringVar(&status, "status", "", "The status to assign to each address.")
	return fs
}
