package list

import (
	"context"
	"flag"
	"fmt"
	_ "log/slog"

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

	// logger := slog.Default()

	subs_db, err := database.NewSubscriptionsDatabase(ctx, opts.SubscriptionsDatabaseURI)

	if err != nil {
		return fmt.Errorf("Failed to instantiate subscriptions database, %w", err)
	}

	defer subs_db.Close()

	subs_cb := func(ctx context.Context, sub *subscription.Subscription) error {
		fmt.Println(sub.Address)
		return nil
	}

	err = subs_db.ListSubscriptions(ctx, subs_cb)

	if err != nil {
		return fmt.Errorf("Failed to list subscriptions, %w", err)
	}

	return nil
}
