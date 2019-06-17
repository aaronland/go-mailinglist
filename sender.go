package mailinglist

import (
	"github.com/aaronland/gomail"
)

func SendMailToList(sender gomail.Sender, subs SubscriptionDatabase, msg *gomail.Message) error {

	cb := func(sub *Subscriber) error {
		msg.SetHeader("To", sub.Address)
		return gomail.Send(sender, msg)
	}

	return subs.ConfirmedSubscriptions(cb)
}
