package status

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"time"

	"github.com/aaronland/go-mailinglist/v2"
	"github.com/aaronland/go-mailinglist/v2/database"
	"github.com/aaronland/go-mailinglist/v2/subscription"
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

	logger := slog.Default()

	subs_db, err := database.NewSubscriptionsDatabase(ctx, opts.SubscriptionsDatabaseURI)

	if err != nil {
		return fmt.Errorf("Failed to instantiate subscriptions database, %w", err)
	}

	defer subs_db.Close()

	var status_id int

	switch opts.Status {
	case "pending":
		status_id = subscription.SUBSCRIPTION_STATUS_PENDING
	case "enabled", "confirmed":
		status_id = subscription.SUBSCRIPTION_STATUS_ENABLED
	case "disabled":
		status_id = subscription.SUBSCRIPTION_STATUS_DISABLED
	case "blocked":
		status_id = subscription.SUBSCRIPTION_STATUS_BLOCKED
	default:
		return fmt.Errorf("Invalid status")
	}

	for _, addr := range opts.Addresses {

		sub, err := subs_db.GetSubscriptionWithAddress(ctx, addr)

		if err != nil {
			return fmt.Errorf("Failed to retrieve subscription for %s, %w", addr, err)
		}

		sub.Status = status_id

		if opts.Status == "confirmed" && sub.Confirmed == 0 {
			now := time.Now()
			sub.Confirmed = now.Unix()
		}

		sub, err = mailinglist.UpdateSubscription(ctx, subs_db, sub)

		if err != nil {
			return fmt.Errorf("Failed to update subscription for %s, %w", addr, err)
		}

		logger.Info("Subscription updated", "address", addr)
	}

	return nil
}
