package errors

import (
	"fmt"
)

type MissingTemplate struct {
	name string
}

func (e *MissingTemplate) Error() string {
	return fmt.Sprintf("Missing template '%s'", e.name)
}

func NewMissingTemplateError(t string) *MissingTemplate {
	return &MissingTemplate{t}
}
