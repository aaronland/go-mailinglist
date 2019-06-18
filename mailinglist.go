package mailinglist

import (
	"context"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/gomail"	
)

func SendMailToList(sender gomail.Sender, subs database.SubscriptionDatabase, msg *gomail.Message) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return SendMailToListWithContext(ctx, sender, subs, msg)
}

func SendMailToListWithContext(ctx context.Context, sender gomail.Sender, subs database.SubscriptionDatabase, msg *gomail.Message) error {

	cb := func(sub *Subscriber) error {
		msg.SetHeader("To", sub.Address)
		return gomail.Send(sender, msg)
	}

	return subs.ConfirmedSubscriptions(ctx, cb)
}

