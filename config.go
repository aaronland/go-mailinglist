package mailinglist

import (
	"net/url"
)

type MailingListConfig struct {
	URL    *url.URL
	Name   string
	Sender string
	Paths  *PathConfig
}

type PathConfig struct {
	Index       string
	Subscribe   string
	Unsubscribe string
	Confirm     string
}
