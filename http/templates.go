package http

import (
	"github.com/aaronland/go-mailinglist/errors"
	"html/template"
	_ "log"
	gohttp "net/http"
)

func LoadTemplate(t *template.Template, name string) (*template.Template, error) {

	named_t := t.Lookup(name)

	if named_t == nil {
		return nil, errors.NewMissingTemplateError(name)
	}

	return named_t, nil
}

func RenderTemplate(rsp gohttp.ResponseWriter, t *template.Template, vars interface{}) {

	err := t.Execute(rsp, vars)

	if err != nil {
		app_err := NewApplicationError(err, E_TEMPLATE_RENDER)
		gohttp.Error(rsp, app_err.Error(), gohttp.StatusInternalServerError)
	}
}
