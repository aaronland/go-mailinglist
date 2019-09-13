package mailinglist

type MailingListConfig struct {
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
