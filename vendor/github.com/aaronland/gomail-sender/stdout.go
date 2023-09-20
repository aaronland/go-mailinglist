package sender

import (
	"context"
	"io"
	"os"

	"github.com/aaronland/gomail/v2"
)

func init() {

	ctx := context.Background()
	err := RegisterSender(ctx, "stdout", NewStdoutSender)

	if err != nil {
		panic(err)
	}
}

// NullSender implements the `gomail.Sender` inferface for delivery messages to STDOUT.
type StdoutSender struct {
	gomail.Sender
}

// NewStdoutSender returns a new `NullSender` instance for delivery messages to STDOUT,
// configured by 'uri' which is expected to take the form of:
//
//	$> stdout://
func NewStdoutSender(ctx context.Context, uri string) (gomail.Sender, error) {

	s := StdoutSender{}
	return &s, nil
}

// Send will copy 'msg' to `os.Stdout`..
func (s *StdoutSender) Send(from string, to []string, msg io.WriterTo) error {

	_, err := msg.WriteTo(os.Stdout)

	if err != nil {
		return err
	}

	return nil
}
