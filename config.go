package mailinglist

import (
	"net/url"
)

type MailingListConfig struct {
	URL          *url.URL
	Name         string
	Sender       string
	Paths        *PathConfig
	FeatureFlags *FeatureFlags
}

type FeatureFlags struct {
	Subscribe     bool
	Unsubscribe   bool
	InviteAccept  bool
	InviteRequest bool
	Confirm       bool
}

type PathConfig struct {
	Index         string
	Subscribe     string
	Unsubscribe   string
	Confirm       string
	InviteRequest string
	InviteAccept  string
}
