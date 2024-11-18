package database

import (
	"context"

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

func (db *DeliveriesDocstoreDatabase) ListDeliveries(ctx context.Context, db ListDeliveriesFunc) error {

}

func (db *DeliveriesDocstoreDatabase) GetDeliveryWithAddressAndMessageId(ctx context.Context, addr string, id string) (*delivery.Delivery, error) {

}

func (db *DeliveriesDocstoreDatabase) Close() error {
	return db.collection.Close()
}
