package message

import (
	"bufio"
	"bytes"
	html_template "html/template"

	"github.com/aaronland/gomail"	
)

func NewMessageFromHTMLTemplate(t *html_template.Template, vars interface{}) (*gomail.Message, error) {

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	err := t.Execute(wr, vars)

	if err != nil {
		return nil, err
	}

	wr.Flush()

	m := gomail.NewMessage()
	m.SetBody("text/html", buf.String())

	return m, nil
}
