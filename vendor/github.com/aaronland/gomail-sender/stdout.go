package sender

import (
	"context"
	"github.com/aaronland/gomail/v2"
	"io"
	_ "log"
	"os"
)

func init() {

	ctx := context.Background()
	err := RegisterSender(ctx, "stdout", NewStdoutSender)

	if err != nil {
		panic(err)
	}
}

type StdoutSender struct {
	gomail.Sender
}

func NewStdoutSender(ctx context.Context, uri string) (gomail.Sender, error) {

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
