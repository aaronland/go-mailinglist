package http

import (
	"html/template"
	_ "log"
	gohttp "net/http"

	"github.com/aaronland/go-mailinglist/v2/errors"
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
