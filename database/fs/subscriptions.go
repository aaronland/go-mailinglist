package fs

import (
	"context"
	"errors"
	"fmt"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/subscription"
	"os"
	"path/filepath"
)

type FSSubscriptionsDatabase struct {
	database.SubscriptionsDatabase
	root string
}

func NewFSSubscriptionsDatabase(root string) (database.SubscriptionsDatabase, error) {

	abs_root, err := filepath.Abs(root)

	if err != nil {
		return nil, err
	}

	info, err := os.Stat(abs_root)

	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, errors.New("Root is not a directory")
	}

	/*
		if info.Mode() != 0700 {
			return nil, errors.New("Root permissions must be 0700")
		}
	*/

	db := FSSubscriptionsDatabase{
		root: abs_root,
	}

	return &db, nil
}

func (db *FSSubscriptionsDatabase) AddSubscription(sub *subscription.Subscription) error {

	path := db.pathForSubscription(sub)

	_, err := os.Stat(path)

	if err == nil {
		return nil
	}

	return db.writeSubscription(sub, path)
}

func (db *FSSubscriptionsDatabase) RemoveSubscription(sub *subscription.Subscription) error {

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

func (db *FSSubscriptionsDatabase) UpdateSubscription(sub *subscription.Subscription) error {

	path := db.pathForSubscription(sub)

	_, err := os.Stat(path)

	if err == nil {
		return nil
	}

	return db.writeSubscription(sub, path)
}

func (db *FSSubscriptionsDatabase) GetSubscriptionWithAddress(addr string) (*subscription.Subscription, error) {

	path := db.pathForAddress(addr)

	_, err := os.Stat(path)

	if err != nil {

		if os.IsNotExist(err) {
			return nil, new(database.NoRecordError)
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

func (db *FSSubscriptionsDatabase) ListSubscriptionsConfirmed(ctx context.Context, cb database.ListSubscriptionsFunc) error {

	local_cb := func(ctx context.Context, sub *subscription.Subscription) error {

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		if !sub.IsConfirmed() {
			return nil
		}

		return cb(sub)
	}

	return db.crawlSubscriptions(ctx, local_cb)
}

func (db *FSSubscriptionsDatabase) ListSubscriptionsUnconfirmed(ctx context.Context, cb database.ListSubscriptionsFunc) error {

	local_cb := func(ctx context.Context, sub *subscription.Subscription) error {

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		if sub.IsConfirmed() {
			return nil
		}

		return cb(sub)
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
	return db.pathForAddress(sub.Address)
}

func (db *FSSubscriptionsDatabase) pathForAddress(addr string) string {
	fname := fmt.Sprintf("%s.json", addr)
	return filepath.Join(db.root, fname)
}
