package deliver

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"log/slog"
	"net/mail"
	"time"

	"github.com/aaronland/go-mailinglist/v2/database"
	"github.com/aaronland/go-mailinglist/v2/delivery"
	"github.com/aaronland/go-mailinglist/v2/eventlog"
	"github.com/aaronland/go-mailinglist/v2/message"
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

	hash := sha256.Sum256([]byte(opts.Body))
	msg_id := hex.EncodeToString(hash[:])

	msg := gomail.NewMessage()
	msg.SetHeader("X-MailingList-Id", msg_id)

	msg.SetBody("text/plain", opts.Body)

	// Something something something... attachment(s)

	deliver_message := func(ctx context.Context, to string) error {

		logger := slog.Default()
		logger = logger.With("to", to)
		logger = logger.With("message id", msg_id)

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

		d, err := deliveries_db.GetDeliveryWithAddressAndMessageId(ctx, to, msg_id)

		if err != nil && !database.IsNotExist(err) {
			event_status = eventlog.EVENTLOG_SEND_FAIL_EVENT
			event_message = err.Error()
			return fmt.Errorf("Failed to retrieve delivery for %s (%s), %w", to, msg_id, err)
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
			MessageId: msg_id,
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

	for _, addr := range opts.To {

		sub, err := subs_db.GetSubscriptionWithAddress(ctx, addr)

		if err != nil {
			return fmt.Errorf("Failed to retrieve subscription for %s, %w", addr, err)
		}

		err = deliver_message(ctx, sub.Address)

		if err != nil {
			return fmt.Errorf("Failed to deliver message to %s, %w", sub.Address, err)
		}
	}

	return nil
}
