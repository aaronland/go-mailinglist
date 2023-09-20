package sender

import (
	"context"
	"io"

	"github.com/aaronland/gomail/v2"
)

func init() {

	ctx := context.Background()
	err := RegisterSender(ctx, "null", NewNullSender)

	if err != nil {
		panic(err)
	}
}

// NullSender implements the `gomail.Sender` inferface for delivery messages to nowhere.
type NullSender struct {
	gomail.Sender
}

// NewNullSender returns a new `NullSender` instance for delivery messages to nowhere,
// configured by 'uri' which is expected to take the form of:
//
//	$> null://
func NewNullSender(ctx context.Context, uri string) (gomail.Sender, error) {

	s := NullSender{}
	return &s, nil
}

// Send will copy 'msg' to a `io.Discard` instance.
func (s *NullSender) Send(from string, to []string, msg io.WriterTo) error {

	_, err := msg.WriteTo(io.Discard)

	if err != nil {
		return err
	}

	return nil
}
