package message

import (
	"context"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/subscription"
	"github.com/aaronland/gomail"
	"log"
	"net/mail"
	"sync"
	"time"
)

type PreSendMessageFilterFunc func(*gomail.Message, *mail.Address) (bool, error) // true to send mail, false to skip

type PostSendMessageFunc func(*gomail.Message, *mail.Address, time.Duration, error) error

type SendMessageOptions struct {
	Sender            gomail.Sender
	Subject           string
	From              *mail.Address
	To                *mail.Address
	PreSendFilterFunc PreSendMessageFilterFunc
	PostSendFunc      PostSendMessageFunc
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

	// please for to be making throttles part of SendMessageOptions - for
	// today we'll just say 10 per second (20190628/thisisaaronland)

	rate := time.Second / 10
	throttle := time.Tick(rate)

	wg := new(sync.WaitGroup)

	cb := func(sub *subscription.Subscription) error {

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		to, err := mail.ParseAddress(sub.Address)

		if err != nil {
			return err
		}

		if opts.PreSendFilterFunc != nil {

			include, err := opts.PreSendFilterFunc(msg, to)

			if err != nil {
				return err
			}

			if !include {
				return nil
			}
		}

		local_opts := &SendMessageOptions{
			Sender:  opts.Sender,
			Subject: opts.Subject,
			From:    opts.From,
			To:      to,
		}

		if opts.PostSendFunc != nil {
			local_opts.PostSendFunc = opts.PostSendFunc
		}

		<-throttle

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		wg.Add(1)

		go func(wg *sync.WaitGroup, msg *gomail.Message, local_opts *SendMessageOptions) {

			t2 := time.Now()

			defer func() {
				wg.Done()
			}()

			err := SendMessage(msg, local_opts)

			if err != nil {
				log.Printf("Failed to send message to %s (%s)\n", to, err)
			}

			if local_opts.PostSendFunc != nil {

				tts := time.Since(t2)

				post_err := local_opts.PostSendFunc(msg, local_opts.To, tts, err)

				if post_err != nil {
					log.Printf("Failed to complete post send message func for %s (%s)\n", to, post_err)
				}
			}

		}(wg, msg, local_opts)

		return nil
	}

	err := subs_db.ListSubscriptionsWithStatus(ctx, cb, subscription.SUBSCRIPTION_STATUS_ENABLED)

	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}
