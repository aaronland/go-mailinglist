package remove

import (
	"context"
	"flag"
	"fmt"
	"log/slog"

	"github.com/aaronland/go-mailinglist/v2/database"
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

		sub, err := subs_db.GetSubscriptionWithAddress(ctx, addr)

		if err != nil {
			return fmt.Errorf("Failed to retrieve subscription for %s, %w", addr, err)
		}

		err = subs_db.RemoveSubscription(ctx, sub)

		if err != nil {
			return fmt.Errorf("Failed to remove subscription for %s, %w", addr, err)
		}

		logger.Info("Subscription removed", "address", addr)
	}

	return nil
}
