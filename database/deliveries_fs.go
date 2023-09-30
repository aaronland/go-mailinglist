package database

import (
	"context"
	"fmt"
	_ "log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/aaronland/go-mailinglist/delivery"
)

type FSDeliveriesDatabase struct {
	DeliveriesDatabase
	root string
}

func init() {
	ctx := context.Background()
	RegisterDeliveriesDatabase(ctx, "fs", NewDeliveriesDatabase)
}

func NewFSDeliveriesDatabase(ctx context.Context, uri string) (DeliveriesDatabase, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	root := u.Path

	abs_root, err := ensureRoot(root)

	if err != nil {
		return nil, err
	}

	db := FSDeliveriesDatabase{
		root: abs_root,
	}

	return &db, nil
}

func (db *FSDeliveriesDatabase) AddDelivery(ctx context.Context, d *delivery.Delivery) error {

	root := filepath.Join(db.root, d.MessageId)

	_, err := os.Stat(root)

	if err != nil {

		if !os.IsNotExist(err) {
			return err
		}

		err = os.MkdirAll(root, 0700)

		if err != nil {
			return err
		}
	}

	fname := fmt.Sprintf("%s.json", d.Address)
	path := filepath.Join(root, fname)

	_, err = os.Stat(path)

	if err == nil {
		return nil
	}

	return marshalData(d, path)
}

func (db *FSDeliveriesDatabase) GetDeliveryWithAddressAndMessageId(ctx context.Context, address string, message_id string) (*delivery.Delivery, error) {

	root := filepath.Join(db.root, message_id)

	_, err := os.Stat(root)

	if err != nil {
		return nil, err
	}

	fname := fmt.Sprintf("%s.json", address)
	path := filepath.Join(root, fname)

	_, err = os.Stat(path)

	if err == nil {
		return nil, err
	}

	return db.readDelivery(path)
}

func (db *FSDeliveriesDatabase) ListDeliveries(ctx context.Context, callback ListDeliveriesFunc) error {

	local_cb := func(ctx context.Context, d *delivery.Delivery) error {

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		return callback(ctx, d)
	}

	return db.crawlDeliveries(ctx, local_cb)
}

func (db *FSDeliveriesDatabase) crawlDeliveries(ctx context.Context, cb func(ctx context.Context, d *delivery.Delivery) error) error {

	local_cb := func(ctx context.Context, path string) error {

		sub, err := db.readDelivery(path)

		if err != nil {
			return err
		}

		return cb(ctx, sub)
	}

	return crawlDatabase(ctx, db.root, local_cb)
}

func (db *FSDeliveriesDatabase) readDelivery(path string) (*delivery.Delivery, error) {

	ev, err := unmarshalData(path, "delivery")

	if err != nil {
		return nil, err
	}

	return ev.(*delivery.Delivery), nil
}
