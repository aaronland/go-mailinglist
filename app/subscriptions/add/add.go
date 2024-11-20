package add

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"time"

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

	for _, addr := range opts.Addresses {

		sub, err := subscription.NewSubscription(addr)

		if err != nil {
			return fmt.Errorf("Failed to add subscription for %s, %w", addr, err)
		}

		if opts.Confirmed {
			now := time.Now()
			sub.Confirmed = now.Unix()
			sub.Status = subscription.SUBSCRIPTION_STATUS_ENABLED
		}

		err = subs_db.AddSubscription(ctx, sub)

		if err != nil {
			return fmt.Errorf("Failed to add subscription for %s (%s), %w", addr, sub.Address, err)
		}

		logger.Info("Subscription added", "address", addr)
	}

	return nil
}
