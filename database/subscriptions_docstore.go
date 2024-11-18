package database

import (
	"context"

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
}

func (db *SubscriptionsDocstoreDatabase) ListSubscriptions(ctx context.Context, cb ListSubscriptionsFunc) error {

}

func (db *SubscriptionsDocstoreDatabase) ListSubscriptionsWithStatus(ctx context.Context, cb ListSubscriptionsFunc, statuses ...int) error {
}

func (db *SubscriptionsDocstoreDatabase) Close() error {
	return db.collection.Close()
}
