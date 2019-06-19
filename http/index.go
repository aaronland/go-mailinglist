package http

import (
	"html/template"
	gohttp "net/http"
)

type IndexHandlerOptions struct {
}

func IndexHandler(opts *IndexHandlerOptions) (gohttp.Handler, error) {

	index_t := template.New("index")

	index_t, err := index_t.Parse(`<html><head><title>Signup</title></head>
<body>
</body></html>`)

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
