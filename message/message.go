package message

import (
	"context"
	"github.com/aaronland/go-mailinglist/database"	
	"github.com/aaronland/go-mailinglist/subscription"
	"github.com/aaronland/gomail"
	"net/mail"
)

type SendMessageOptions struct {
	Sender gomail.Sender
	Subject string	
	From   *mail.Address	
	To     *mail.Address
}

func SendMessage(msg *gomail.Message, opts *SendMessageOptions) error {

	msg.SetHeader("From", opts.From.Address)
	msg.SetHeader("To", opts.To.Address)
	msg.SetHeader("Subject", opts.Subject)
	
	return gomail.Send(opts.Sender, msg)
}

func SendMessageToList(subs_db database.SubscriptionsDatabase, msg *gomail.Message, opts *SendMessageOptions) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return SendMailToListWithContext(ctx, subs_db, msg, opts)
}

func SendMailToListWithContext(ctx context.Context, subs_db database.SubscriptionsDatabase, msg *gomail.Message, opts *SendMessageOptions) error {

	cb := func(sub *subscription.Subscription) error {

		// throttles and goroutines and stuff
		
		to, err := mail.ParseAddress(sub.Address)

		if err != nil {
			return err
		}

		local_opts := &SendMessageOptions{
			Sender: opts.Sender,
			Subject: opts.Subject,			
			From:   opts.From,
			To:     to,
		}

		return SendMessage(msg, local_opts)
	}

	return subs_db.ListSubscriptionsConfirmed(ctx, cb)
}
