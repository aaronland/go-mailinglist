package message

import (
	"context"
	"log"
	"net/mail"
	"sync"
	"time"

	"github.com/aaronland/go-mailinglist/v2/database"
	"github.com/aaronland/go-mailinglist/v2/subscription"
	"github.com/aaronland/gomail/v2"
)

type PreSendMessageFilterFunc func(context.Context, *gomail.Message, *mail.Address) (bool, error) // true to send mail, false to skip

type PostSendMessageFunc func(context.Context, *gomail.Message, *mail.Address, time.Duration, error) error

type SendMessageOptions struct {
	Sender            gomail.Sender
	Subject           string
	From              *mail.Address
	To                *mail.Address
	PreSendFilterFunc PreSendMessageFilterFunc
	PostSendFunc      PostSendMessageFunc
	// Throttle	<-chan time.Time
}

func SendMessage(ctx context.Context, opts *SendMessageOptions, msg *gomail.Message) error {

	from := opts.From.String()
	to := opts.To.String()

	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", opts.Subject)

	return gomail.Send(opts.Sender, msg)
}

func SendMessageToList(ctx context.Context, subs_db database.SubscriptionsDatabase, msg *gomail.Message, opts *SendMessageOptions) error {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	t1 := time.Now()

	defer func() {
		log.Printf("Time to send message to list %v\n", time.Since(t1))
	}()

	// please for to be making throttles part of SendMessageOptions - for
	// today we'll just say 10 per second (20190628/thisisaaronland)

	rate := time.Second / 10
	throttle := time.Tick(rate)

	wg := new(sync.WaitGroup)

	cb := func(ctx context.Context, sub *subscription.Subscription) error {

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

			include, err := opts.PreSendFilterFunc(ctx, msg, to)

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

		go func(wg *sync.WaitGroup, local_opts *SendMessageOptions, msg *gomail.Message) {

			t2 := time.Now()

			defer func() {
				wg.Done()
			}()

			err := SendMessage(ctx, local_opts, msg)

			if err != nil {
				log.Printf("Failed to send message to %s (%s)\n", to, err)
			}

			if local_opts.PostSendFunc != nil {

				tts := time.Since(t2)

				post_err := local_opts.PostSendFunc(ctx, msg, local_opts.To, tts, err)

				if post_err != nil {
					log.Printf("Failed to complete post send message func for %s (%s)\n", to, post_err)
				}
			}

		}(wg, local_opts, msg)

		return nil
	}

	err := subs_db.ListSubscriptionsWithStatus(ctx, subscription.SUBSCRIPTION_STATUS_ENABLED, cb)

	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}
