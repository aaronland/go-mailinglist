package message

import (
	"context"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/subscription"
	"github.com/aaronland/gomail"
	"log"
	"net/mail"
	"time"
)

type SendMessageOptions struct {
	Sender  gomail.Sender
	Subject string
	From    *mail.Address
	To      *mail.Address
	// Throttle	<-chan time.Time
}

func SendMessage(msg *gomail.Message, opts *SendMessageOptions) error {

	from := opts.From.String()
	to := opts.To.String()

	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", opts.Subject)

	return gomail.Send(opts.Sender, msg)
}

func SendMessageToList(subs_db database.SubscriptionsDatabase, msg *gomail.Message, opts *SendMessageOptions) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return SendMailToListWithContext(ctx, subs_db, msg, opts)
}

func SendMailToListWithContext(ctx context.Context, subs_db database.SubscriptionsDatabase, msg *gomail.Message, opts *SendMessageOptions) error {

	t1 := time.Now()

	defer func() {
		log.Printf("Time to send message to list %v\n", time.Since(t1))
	}()

	// please for to be making this part of SendMessageOptions
	// for today we'll just say 10 per second
	// (20190628/thisisaaronland)

	rate := time.Second / 10
	throttle := time.Tick(rate)

	cb := func(sub *subscription.Subscription) error {

		to, err := mail.ParseAddress(sub.Address)

		if err != nil {
			return err
		}

		local_opts := &SendMessageOptions{
			Sender:  opts.Sender,
			Subject: opts.Subject,
			From:    opts.From,
			To:      to,
		}

		<-throttle

		go func(msg *gomail.Message, local_opts *SendMessageOptions) {

			t1 := time.Now()

			defer func() {
				log.Printf("Time to send message to %s %v\n", local_opts.To, time.Since(t1))
			}()

			err := SendMessage(msg, local_opts)

			if err != nil {
				log.Printf("Failed to send message to %s (%s)\n", to, err)
			}

		}(msg, local_opts)

		return nil
	}

	return subs_db.ListSubscriptionsWithStatus(ctx, cb, subscription.SUBSCRIPTION_STATUS_ENABLED)
}
