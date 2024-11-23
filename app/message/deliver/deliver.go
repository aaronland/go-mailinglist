package deliver

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/mail"
	"time"

	"github.com/aaronland/go-mailinglist/v2/database"
	"github.com/aaronland/go-mailinglist/v2/delivery"
	"github.com/aaronland/go-mailinglist/v2/eventlog"
	"github.com/aaronland/go-mailinglist/v2/message"
	"github.com/aaronland/go-mailinglist/v2/subscription"
	"github.com/aaronland/gocloud-blob/bucket"
	"github.com/aaronland/gomail-sender"
	"github.com/aaronland/gomail/v2"
)

func Run(ctx context.Context) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	opts, err := RunOptionsFromFlagSet(fs)

	if err != nil {
		return fmt.Errorf("Failed to derive options from flagset, %w", err)
	}

	return RunWithOptions(ctx, opts)
}

func RunWithOptions(ctx context.Context, opts *RunOptions) error {

	subs_db, err := database.NewSubscriptionsDatabase(ctx, opts.SubscriptionsDatabaseURI)

	if err != nil {
		return fmt.Errorf("Failed to instantiate subscriptions database, %w", err)
	}

	defer subs_db.Close()

	deliveries_db, err := database.NewDeliveriesDatabase(ctx, opts.DeliveriesDatabaseURI)

	if err != nil {
		return fmt.Errorf("Failed to instantiate deliveries database, %w", err)
	}

	defer deliveries_db.Close()

	eventlogs_db, err := database.NewEventLogsDatabase(ctx, opts.EventLogsDatabaseURI)

	if err != nil {
		return fmt.Errorf("Failed to instantiate event logs database, %w", err)
	}

	defer eventlogs_db.Close()

	message_sender, err := sender.NewSender(ctx, opts.SenderURI)

	if err != nil {
		return fmt.Errorf("Failed to create sender, %w", err)
	}

	from_addr, err := mail.ParseAddress(opts.From)

	if err != nil {
		return fmt.Errorf("Invalid from address, %w", err)
	}

	// https://pkg.go.dev/github.com/aaronland/gomail/v2#Message

	msg := gomail.NewMessage()
	msg.SetHeader("X-MailingList-Id", opts.MessageId)

	msg.SetBody(opts.ContentType, opts.Body)

	for _, uri := range opts.Attachments {

		bucket_uri, bucket_key, err := bucket.ParseURI(uri)

		if err != nil {
			return fmt.Errorf("Failed to derive bucket URI and key from attachment, %w", err)
		}

		b, err := bucket.OpenBucket(ctx, bucket_uri)

		if err != nil {
			return err
		}

		defer b.Close()

		r, err := b.NewReader(ctx, bucket_key, nil)

		if err != nil {
			return fmt.Errorf("Failed to open attachement (%s), %w", uri, err)
		}

		defer r.Close()

		msg.EmbedReader(uri, r)
	}

	deliver_message := func(ctx context.Context, to string) error {

		logger := slog.Default()
		logger = logger.With("to", to)
		logger = logger.With("message id", opts.MessageId)

		to_addr, err := mail.ParseAddress(to)

		if err != nil {
			return fmt.Errorf("Invalid to address (%s), %w", to, err)
		}

		event_status := eventlog.EVENTLOG_CUSTOM_EVENT
		event_message := ""

		defer func() {

			now := time.Now()

			ev := &eventlog.EventLog{
				Address: to,
				Created: now.Unix(),
				Event:   event_status,
				Message: event_message,
			}

			err := eventlogs_db.AddEventLog(ctx, ev)

			if err != nil {
				logger.Error("Failed to add event log", "status", event_status, "message", event_message, "error", err)
			}
		}()

		d, err := deliveries_db.GetDeliveryWithAddressAndMessageId(ctx, to, opts.MessageId)

		if err != nil && !database.IsNotExist(err) {
			event_status = eventlog.EVENTLOG_SEND_FAIL_EVENT
			event_message = err.Error()
			return fmt.Errorf("Failed to retrieve delivery for %s (%s), %w", to, opts.MessageId, err)
		}

		if d != nil {
			event_status = eventlog.EVENTLOG_SEND_DUPLICATE_EVENT
			logger.Info("Message already delivered")
			return nil
		}

		send_opts := &message.SendMessageOptions{
			Sender:  message_sender,
			Subject: opts.Subject,
			From:    from_addr,
			To:      to_addr,
		}

		err = message.SendMessage(ctx, send_opts, msg)

		if err != nil {
			event_status = eventlog.EVENTLOG_SEND_FAIL_EVENT
			event_message = err.Error()
			return fmt.Errorf("Failed to deliver message to %s, %w", err)
		}

		now := time.Now()

		new_d := &delivery.Delivery{
			MessageId: opts.MessageId,
			Address:   to,
			Delivered: now.Unix(),
		}

		err = deliveries_db.AddDelivery(ctx, new_d)

		if err != nil {
			logger.Error("Message delivered but failed to add to deliveries database", "error", err)
		}

		event_status = eventlog.EVENTLOG_SEND_OK_EVENT
		return err
	}

	switch len(opts.To) {
	case 0:

		subs_cb := func(ctx context.Context, sub *subscription.Subscription) error {

			err = deliver_message(ctx, sub.Address)

			if err != nil {
				return fmt.Errorf("Failed to deliver message to %s, %w", sub.Address, err)
			}

			return nil
		}

		err := subs_db.ListSubscriptionsWithStatus(ctx, subscription.SUBSCRIPTION_STATUS_ENABLED, subs_cb)

		if err != nil {
			return fmt.Errorf("Failed to list subscriptions, %w", err)
		}

	default:

		for _, addr := range opts.To {

			sub, err := subs_db.GetSubscriptionWithAddress(ctx, addr)

			if err != nil {
				return fmt.Errorf("Failed to retrieve subscription for %s, %w", addr, err)
			}

			if sub.Status != subscription.SUBSCRIPTION_STATUS_ENABLED {
				continue
			}
			
			err = deliver_message(ctx, sub.Address)

			if err != nil {
				return fmt.Errorf("Failed to deliver message to %s, %w", sub.Address, err)
			}
		}
	}

	return nil
}
