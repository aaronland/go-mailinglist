package sender

import (
	"context"
	"github.com/aaronland/gomail/v2"
	"io"
	_ "log"
)

func init() {

	ctx := context.Background()
	err := RegisterSender(ctx, "null", NewNullSender)

	if err != nil {
		panic(err)
	}
}

type NullSender struct {
	gomail.Sender
}

func NewNullSender(ctx context.Context, uri string) (gomail.Sender, error) {

	s := NullSender{}
	return &s, nil
}

func (s *NullSender) Send(from string, to []string, msg io.WriterTo) error {

	_, err := msg.WriteTo(io.Discard)

	if err != nil {
		return err
	}

	return nil
}
