package database

import (
	"context"
	"fmt"
	"io"
	_ "log/slog"

	"github.com/aaronland/go-mailinglist/v2/subscription"
	aa_docstore "github.com/aaronland/gocloud-docstore"
	"gocloud.dev/docstore"
)

type SubscriptionsDocstoreDatabase struct {
	SubscriptionsDatabase
	collection *docstore.Collection
}

func init() {

	ctx := context.Background()

	err := RegisterSubscriptionsDatabase(ctx, "awsdynamodb", NewSubscriptionsDocstoreDatabase)

	if err != nil {
		panic(err)
	}

	for _, scheme := range docstore.DefaultURLMux().CollectionSchemes() {

		err := RegisterSubscriptionsDatabase(ctx, scheme, NewSubscriptionsDocstoreDatabase)

		if err != nil {
			panic(err)
		}

	}
}

func NewSubscriptionsDocstoreDatabase(ctx context.Context, uri string) (SubscriptionsDatabase, error) {

	col, err := aa_docstore.OpenCollection(ctx, uri)

	if err != nil {
		return nil, err
	}

	db := &SubscriptionsDocstoreDatabase{
		collection: col,
	}

	return db, nil
}

func (db *SubscriptionsDocstoreDatabase) AddSubscription(ctx context.Context, sub *subscription.Subscription) error {
	return db.collection.Put(ctx, sub)
}

func (db *SubscriptionsDocstoreDatabase) RemoveSubscription(ctx context.Context, sub *subscription.Subscription) error {
	return db.collection.Delete(ctx, sub)
}

func (db *SubscriptionsDocstoreDatabase) UpdateSubscription(ctx context.Context, sub *subscription.Subscription) error {
	return db.collection.Replace(ctx, sub)
}

func (db *SubscriptionsDocstoreDatabase) GetSubscriptionWithAddress(ctx context.Context, addr string) (*subscription.Subscription, error) {
	q := db.collection.Query()
	q = q.Where("address", "=", addr)
	return db.getSubscriptionWithQuery(ctx, q)
}

func (db *SubscriptionsDocstoreDatabase) ListSubscriptions(ctx context.Context, cb ListSubscriptionsFunc) error {
	q := db.collection.Query()
	return db.getSubscriptionsWithCallback(ctx, q, cb)
}

func (db *SubscriptionsDocstoreDatabase) ListSubscriptionsWithStatus(ctx context.Context, status int, cb ListSubscriptionsFunc) error {
	q := db.collection.Query()
	q = q.Where("status", "=", status)
	return db.getSubscriptionsWithCallback(ctx, q, cb)
}

func (db *SubscriptionsDocstoreDatabase) Close() error {
	return db.collection.Close()
}

func (db *SubscriptionsDocstoreDatabase) getSubscriptionWithQuery(ctx context.Context, q *docstore.Query) (*subscription.Subscription, error) {

	iter := q.Get(ctx)
	defer iter.Stop()

	var s subscription.Subscription
	err := iter.Next(ctx, &s)

	if err == io.EOF {
		return nil, NoRecordError("")
	} else if err != nil {
		return nil, fmt.Errorf("Failed to interate, %w", err)
	} else {
		return &s, nil
	}
}

func (db *SubscriptionsDocstoreDatabase) getSubscriptionsWithCallback(ctx context.Context, q *docstore.Query, cb ListSubscriptionsFunc) error {

	// Wut wut wut...
	// https://github.com/google/go-cloud/issues/2426
	// https://pkg.go.dev/gocloud.dev/docstore#FieldPath
	// This only ever yields {"address":"","confirmed":0,"lastmodified":0}
	// if field paths are "address,created,..."

	iter := q.Get(ctx)
	defer iter.Stop()

	for {

		var s subscription.Subscription
		err := iter.Next(ctx, &s)

		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("Failed to interate, %w", err)
		} else {

			err := cb(ctx, &s)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
