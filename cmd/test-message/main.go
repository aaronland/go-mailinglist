package main

import (
	"context"
	"log"
	"net/mail"

	"github.com/aaronland/go-mailinglist/message"
	"github.com/aaronland/gomail-sender"
	"github.com/aaronland/gomail/v2"
	
)

func main() {

	ctx := context.Background()

	s, err := sender.NewSender(ctx, "stdout://")

	if err != nil {
		log.Fatal(err)
	}

	to, _ := mail.ParseAddress("to@example.com")
	from, _ := mail.ParseAddress("from@example.com")
	subject := "This is the subject"

	opts := &message.SendMessageOptions{
		Sender:  s,
		Subject: subject,
		From:    from,
		To:      to,
	}

	m := gomail.NewMessage()
	m.SetBody("text/html", "<p>hello world</p>")

	err = message.SendMessage(m, opts)

	if err != nil {
		log.Fatal(err)
	}

}
