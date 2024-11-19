package www

// Bootstrap stuff is handled by go-http-bootstrap middleware
// see cmd/subscriptiond/main.go for details

import (
	"html/template"
	"net/http"

	"github.com/aaronland/go-mailinglist/v2"
)

type IndexTemplateVars struct {
	SiteName string
	Paths    *mailinglist.PathConfig
	Flags    *mailinglist.FeatureFlags
}

type IndexHandlerOptions struct {
	Config    *mailinglist.MailingListConfig
	Templates *template.Template
}

func IndexHandler(opts *IndexHandlerOptions) (http.Handler, error) {

	index_t, err := LoadTemplate(opts.Templates, "index")

	if err != nil {
		return nil, err
	}

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		vars := IndexTemplateVars{
			SiteName: opts.Config.Name,
			Paths:    opts.Config.Paths,
			Flags:    opts.Config.FeatureFlags,
		}

		RenderTemplate(rsp, index_t, vars)
		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
