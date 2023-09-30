package http

// CSRF crumbs are handled by go-http-crumb middleware
// this package defines a custom go-http-crumb error handler
// see cmd/subscriptiond/main.go for details

import (
	"fmt"
	"html/template"
	gohttp "net/http"

	_ "github.com/aaronland/go-http-crumb"
	"github.com/aaronland/go-mailinglist"
)

type CrumbErrorHandlerOptions struct {
	Config    *mailinglist.MailingListConfig
	Templates *template.Template
}

type CrumbErrorVars struct {
	SiteName string
	Paths    *mailinglist.PathConfig
	Error    error
}

func CrumbErrorHandler(opts *CrumbErrorHandlerOptions) (gohttp.Handler, error) {

	error_t, err := LoadTemplate(opts.Templates, "crumb_error")

	if err != nil {
		return nil, err
	}

	handler_fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		// Fix me. Where did this go?
		// crumb_err, _, err := crumb.GetErrorContextValuesWithRequest(req)

		crumb_err := fmt.Errorf("Internal server error")

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		vars := CrumbErrorVars{
			SiteName: opts.Config.Name,
			Paths:    opts.Config.Paths,
			Error:    crumb_err,
		}

		RenderTemplate(rsp, error_t, vars)
		return
	}

	return gohttp.HandlerFunc(handler_fn), nil
}
