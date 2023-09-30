package database

import (
	"context"
	_ "log"
	"net/url"
	"os"

	"github.com/aaronland/go-mailinglist/subscription"
)

type FSSubscriptionsDatabase struct {
	SubscriptionsDatabase
	root string
}

func init() {
	ctx := context.Background()
	RegisterSubscriptionsDatabase(ctx, "fs", NewFSSubscriptionsDatabase)
}

func NewFSSubscriptionsDatabase(ctx context.Context, uri string) (SubscriptionsDatabase, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	root := u.Path

	abs_root, err := ensureRoot(root)

	if err != nil {
		return nil, err
	}

	db := FSSubscriptionsDatabase{
		root: abs_root,
	}

	return &db, nil
}

func (db *FSSubscriptionsDatabase) AddSubscription(ctx context.Context, sub *subscription.Subscription) error {

	path := db.pathForSubscription(sub)

	_, err := os.Stat(path)

	if err == nil {
		return nil
	}

	return db.writeSubscription(sub, path)
}

func (db *FSSubscriptionsDatabase) RemoveSubscription(ctx context.Context, sub *subscription.Subscription) error {

	path := db.pathForSubscription(sub)

	_, err := os.Stat(path)

	if err != nil {

		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	return os.Remove(path)
}

func (db *FSSubscriptionsDatabase) UpdateSubscription(ctx context.Context, sub *subscription.Subscription) error {

	path := db.pathForSubscription(sub)

	_, err := os.Stat(path)

	if err != nil {
		return err
	}

	return db.writeSubscription(sub, path)
}

func (db *FSSubscriptionsDatabase) GetSubscriptionWithAddress(ctx context.Context, addr string) (*subscription.Subscription, error) {

	path := pathForAddress(db.root, addr)

	_, err := os.Stat(path)

	if err != nil {

		if os.IsNotExist(err) {
			return nil, new(NoRecordError)
		}

		return nil, err
	}

	return db.readSubscription(path)
}

func (db *FSSubscriptionsDatabase) readSubscription(path string) (*subscription.Subscription, error) {

	sub, err := unmarshalData(path, "subscription")

	if err != nil {
		return nil, err
	}

	return sub.(*subscription.Subscription), nil
}

func (db *FSSubscriptionsDatabase) writeSubscription(sub *subscription.Subscription, path string) error {

	return marshalData(sub, path)
}

func (db *FSSubscriptionsDatabase) ListSubscriptions(ctx context.Context, cb ListSubscriptionsFunc) error {

	local_cb := func(ctx context.Context, sub *subscription.Subscription) error {

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		return cb(ctx, sub)
	}

	return db.crawlSubscriptions(ctx, local_cb)
}

func (db *FSSubscriptionsDatabase) ListSubscriptionsWithStatus(ctx context.Context, cb ListSubscriptionsFunc, status ...int) error {

	local_cb := func(ctx context.Context, sub *subscription.Subscription) error {

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		match := false

		for _, s := range status {
			if sub.Status == s {
				match = true
				break
			}
		}

		if !match {
			return nil
		}

		return cb(ctx, sub)
	}

	return db.crawlSubscriptions(ctx, local_cb)
}

func (db *FSSubscriptionsDatabase) crawlSubscriptions(ctx context.Context, cb func(ctx context.Context, sub *subscription.Subscription) error) error {

	local_cb := func(ctx context.Context, path string) error {

		sub, err := db.readSubscription(path)

		if err != nil {
			return err
		}

		return cb(ctx, sub)
	}

	return crawlDatabase(ctx, db.root, local_cb)
}

func (db *FSSubscriptionsDatabase) pathForSubscription(sub *subscription.Subscription) string {
	return pathForAddress(db.root, sub.Address)
}
