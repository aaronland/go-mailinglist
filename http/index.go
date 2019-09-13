package http

import (
	"github.com/aaronland/go-mailinglist"
	"html/template"
	gohttp "net/http"
)

type IndexTemplateVars struct {
	Paths *mailinglist.PathConfig
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
			Paths: opts.Config.Paths,
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
