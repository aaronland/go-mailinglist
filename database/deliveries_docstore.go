package database

import (
	"context"
	"fmt"
	"io"

	"github.com/aaronland/go-mailinglist/v2/delivery"
	aa_docstore "github.com/aaronland/gocloud-docstore"
	"gocloud.dev/docstore"
)

type DeliveriesDocstoreDatabase struct {
	DeliveriesDatabase
	collection *docstore.Collection
}

func init() {

	ctx := context.Background()

	err := RegisterDeliveriesDatabase(ctx, "awsdynamodb", NewDeliveriesDocstoreDatabase)

	if err != nil {
		panic(err)
	}

	for _, scheme := range docstore.DefaultURLMux().CollectionSchemes() {

		err := RegisterDeliveriesDatabase(ctx, scheme, NewDeliveriesDocstoreDatabase)

		if err != nil {
			panic(err)
		}

	}

}

func NewDeliveriesDocstoreDatabase(ctx context.Context, uri string) (DeliveriesDatabase, error) {

	col, err := aa_docstore.OpenCollection(ctx, uri)

	if err != nil {
		return nil, err
	}

	db := &DeliveriesDocstoreDatabase{
		collection: col,
	}

	return db, nil
}

func (db *DeliveriesDocstoreDatabase) AddDelivery(ctx context.Context, d *delivery.Delivery) error {
	return db.collection.Put(ctx, d)
}

func (db *DeliveriesDocstoreDatabase) ListDeliveries(ctx context.Context, cb ListDeliveriesFunc) error {
	q := db.collection.Query()
	return db.getDeliveriesWithCallback(ctx, q, cb)
}

func (db *DeliveriesDocstoreDatabase) GetDeliveryWithAddressAndMessageId(ctx context.Context, addr string, msg_id string) (*delivery.Delivery, error) {
	q := db.collection.Query()
	q.Where("MessageId", "=", msg_id)
	q.Where("Address", "=", addr)
	return db.getDeliveryWithQuery(ctx, q)
}

func (db *DeliveriesDocstoreDatabase) Close() error {
	return db.collection.Close()
}

func (db *DeliveriesDocstoreDatabase) getDeliveryWithQuery(ctx context.Context, q *docstore.Query) (*delivery.Delivery, error) {

	iter := q.Get(ctx)
	defer iter.Stop()

	var d delivery.Delivery
	err := iter.Next(ctx, &d)

	if err == io.EOF {
		return nil, NoRecordError("")
	} else if err != nil {
		return nil, fmt.Errorf("Failed to iterate, %w", err)
	} else {
		return &d, nil
	}
}

func (db *DeliveriesDocstoreDatabase) getDeliveriesWithCallback(ctx context.Context, q *docstore.Query, cb ListDeliveriesFunc) error {

	iter := q.Get(ctx)
	defer iter.Stop()

	for {

		var d delivery.Delivery
		err := iter.Next(ctx, &d)

		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("Failed to iterate, %w", err)
		} else {

			err := cb(ctx, &d)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
