package mailinglist

import (
	"context"
	"time"

	"github.com/aaronland/go-mailinglist/v2/database"
	"github.com/aaronland/go-mailinglist/v2/subscription"
)

func UpdateSubscription(ctx context.Context, subs_db database.SubscriptionsDatabase, sub *subscription.Subscription) (*subscription.Subscription, error) {

	now := time.Now()
	ts := now.Unix()

	sub.LastModified = ts

	err := subs_db.UpdateSubscription(ctx, sub)

	if err != nil {
		return sub, err
	}

	return sub, nil
}
