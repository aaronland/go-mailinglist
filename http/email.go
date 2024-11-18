package http

import (
	"github.com/aaronland/go-mailinglist/v2"
)

type ConfirmationEmailTemplateVars struct {
	SiteName string
	SiteRoot string
	Code     string
	URL      string
	Paths    *mailinglist.PathConfig
	Action   string
}
