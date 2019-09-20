package http

// Bootstrap stuff is handled by go-http-bootstrap middleware
// see cmd/subscriptiond/main.go for details

import (
	"github.com/aaronland/go-mailinglist"
	"html/template"
	gohttp "net/http"
)

type IndexTemplateVars struct {
	SiteName string
	Paths    *mailinglist.PathConfig
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
		}

		err := index_t.Execute(rsp, vars)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
		}

		return
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
