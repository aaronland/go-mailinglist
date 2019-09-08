package http

import (
	"html/template"
	gohttp "net/http"
)

type IndexHandlerOptions struct {
	Templates *template.Template
}

func IndexHandler(opts *IndexHandlerOptions) (gohttp.Handler, error) {

	index_t, err := LoadTemplate(opts.Templates, "index")

	if err != nil {
		return nil, err
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		err := index_t.Execute(rsp, nil)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
		}

		return
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
