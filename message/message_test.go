package message

import (
	"context"
	"net/mail"
	"testing"

	"github.com/aaronland/gomail-sender"
	"github.com/aaronland/gomail/v2"
)

func TestSendMessages(t *testing.T) {

	ctx := context.Background()
	s, err := sender.NewStdoutSender(ctx, "stdout://")

	if err != nil {
		t.Fatalf("Failed to create sender, %v", err)
	}

	to, _ := mail.ParseAddress("to@example.com")
	from, _ := mail.ParseAddress("from@example.com")
	subject := "This is the subject"

	opts := &SendMessageOptions{
		Sender:  s,
		Subject: subject,
		From:    from,
		To:      to,
	}

	m := gomail.NewMessage()
	m.SetBody("text/html", "<p>hello world</p>")

	err = SendMessage(ctx, opts, m)

	if err != nil {
		t.Fatalf("Failed to send message, %v", err)
	}

}
