package http

import (
	"github.com/aaronland/go-http-crumb"
	"github.com/aaronland/go-mailinglist"
	"html/template"
	gohttp "net/http"
)

type CrumbErrorHandlerOptions struct {
	Config    *mailinglist.MailingListConfig
	Templates *template.Template
}

type CrumbErrorVars struct {
	Error error
}

func CrumbErrorHandlerFunc(opts *CrumbErrorHandlerOptions) (crumb.ErrorHandlerFunc, error) {

	error_t, err := LoadTemplate(opts.Templates, "crumb_error")

	if err != nil {
		return nil, err
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request, err error, http_status int) gohttp.Handler {

		handler_fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

			vars := CrumbErrorVars{
				Error: err,
			}

			RenderTemplate(rsp, error_t, vars)
			return
		}

		return gohttp.HandlerFunc(handler_fn)
	}

	return fn, nil
}
