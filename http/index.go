package http

import (
	"html/template"
	gohttp "net/http"
)

type IndexTemplateVars struct {
	Paths *PathOptions
}

type IndexHandlerOptions struct {
	Templates *template.Template
	Paths     *PathOptions
}

func IndexHandler(opts *IndexHandlerOptions) (gohttp.Handler, error) {

	index_t, err := LoadTemplate(opts.Templates, "index")

	if err != nil {
		return nil, err
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		vars := IndexTemplateVars{
			Paths: opts.Paths,
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
