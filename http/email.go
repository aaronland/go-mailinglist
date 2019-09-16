package http

import (
	"github.com/aaronland/go-mailinglist"
)

type ConfirmationEmailTemplateVars struct {
	SiteName string
	Code     string
	URL      string
	Paths    *mailinglist.PathConfig
	Action   string
}
