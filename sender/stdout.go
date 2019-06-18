package sender

import (
	"github.com/aaronland/gomail"
	"io"
	"os"
)

type StdoutSender struct {
	gomail.Sender
}

func NewStdoutSender() (gomail.Sender, error) {

	s := StdoutSender{}
	return &s, nil
}

func (s *StdoutSender) Send(from string, to []string, msg io.WriterTo) error {

	_, err := msg.WriteTo(os.Stdout)

	if err != nil {
		return err
	}

	return nil
}
