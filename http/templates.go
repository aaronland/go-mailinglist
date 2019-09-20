package http

import (
	"github.com/aaronland/go-mailinglist/errors"
	"html/template"
	gohttp "net/http"
	"log"
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
		log.Printf("RENDER ERROR '%s'\n", err)		
		gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
	}
}
