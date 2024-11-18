package http

// Bootstrap stuff is handled by go-http-bootstrap middleware
// see cmd/subscriptiond/main.go for details

import (
	"html/template"
	gohttp "net/http"

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

func IndexHandler(opts *IndexHandlerOptions) (gohttp.Handler, error) {

	index_t, err := LoadTemplate(opts.Templates, "index")

	if err != nil {
		return nil, err
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		vars := IndexTemplateVars{
			SiteName: opts.Config.Name,
			Paths:    opts.Config.Paths,
			Flags:    opts.Config.FeatureFlags,
		}

		RenderTemplate(rsp, index_t, vars)
		return
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
