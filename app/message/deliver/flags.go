package deliver

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var subscriptions_database_uri string
var deliveries_database_uri string
var eventlogs_database_uri string

var sender_uri string

var to multi.MultiString
var from string

var subject string
var body string
var content_type string
var message_id string

var attachments multi.MultiString

var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("deliver")

	fs.StringVar(&subscriptions_database_uri, "subscriptions-database-uri", "", "A registered aaronland/go-mailinglist/v2/database.SubscriptionsDatabase URI.")
	fs.StringVar(&deliveries_database_uri, "deliveries-database-uri", "", "A registered aaronland/go-mailinglist/v2/database.DeliveriesDatabase URI.")
	fs.StringVar(&eventlogs_database_uri, "eventlogs-database-uri", "", "A registered aaronland/go-mailinglist/v2/database.EventLogsDatabase URI.")
	fs.StringVar(&sender_uri, "sender-uri", "", "A registered aaronland/go-mail.Sender URI.")

	fs.Var(&to, "to", "One or more addresses to deliver the message to. If empty then the message will be delivered to all subscribers whose status is \"enabled\".")
	fs.StringVar(&from, "from", "", "The address delivering the message.")
	fs.StringVar(&subject, "subject", "", "The subject of the message being delivered.")
	fs.StringVar(&body, "body", "", "The body of the message being delivered. If \"-\" then body will be read from STDIN.")
	fs.StringVar(&content_type, "content-type", "text/plain", "The content-type of the message body.")
	fs.StringVar(&message_id, "message-id", "", "Optional custom message ID to assign to the message. If empty a unique key will be generated on delivery.")

	fs.Var(&attachments, "attachment", "Zero or more URIs referencing files to attach to the message. URIs are dereferenced using the aaronland/gocloud-blob/bucket.ParseURI method.")

	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")
	return fs
}
