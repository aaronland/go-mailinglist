package database

import (
	"context"
	"fmt"
	"io"

	"github.com/aaronland/go-mailinglist/v2/confirmation"
	aa_docstore "github.com/aaronland/gocloud-docstore"
	"gocloud.dev/docstore"
)

type ConfirmationsDocstoreDatabase struct {
	ConfirmationsDatabase
	collection *docstore.Collection
}

func init() {

	ctx := context.Background()

	err := RegisterConfirmationsDatabase(ctx, "awsdynamodb", NewConfirmationsDocstoreDatabase)

	if err != nil {
		panic(err)
	}

	for _, scheme := range docstore.DefaultURLMux().CollectionSchemes() {

		err := RegisterConfirmationsDatabase(ctx, scheme, NewConfirmationsDocstoreDatabase)

		if err != nil {
			panic(err)
		}

	}

}

func NewConfirmationsDocstoreDatabase(ctx context.Context, uri string) (ConfirmationsDatabase, error) {

	col, err := aa_docstore.OpenCollection(ctx, uri)

	if err != nil {
		return nil, err
	}

	db := &ConfirmationsDocstoreDatabase{
		collection: col,
	}

	return db, nil
}

func (db *ConfirmationsDocstoreDatabase) AddConfirmation(ctx context.Context, conf *confirmation.Confirmation) error {
	return db.collection.Put(ctx, conf)
}

func (db *ConfirmationsDocstoreDatabase) RemoveConfirmation(ctx context.Context, conf *confirmation.Confirmation) error {
	return db.collection.Delete(ctx, conf)
}

func (db *ConfirmationsDocstoreDatabase) GetConfirmationWithCode(ctx context.Context, code string) (*confirmation.Confirmation, error) {

	q := db.collection.Query()
	q = q.Where("Code", "=", code)

	return db.getConfirmationWithQuery(ctx, q)
}

func (db *ConfirmationsDocstoreDatabase) ListConfirmations(ctx context.Context, cb ListConfirmationsFunc) error {

	q := db.collection.Query()
	return db.getConfirmationsWithCallback(ctx, q, cb)
}

func (db *ConfirmationsDocstoreDatabase) Close() error {
	return db.collection.Close()
}

func (db *ConfirmationsDocstoreDatabase) getConfirmationWithQuery(ctx context.Context, q *docstore.Query) (*confirmation.Confirmation, error) {

	iter := q.Get(ctx)
	defer iter.Stop()

	var c confirmation.Confirmation
	err := iter.Next(ctx, &c)

	if err == io.EOF {
		return nil, NoRecordError("")
	} else if err != nil {
		return nil, fmt.Errorf("Failed to interate, %w", err)
	} else {
		return &c, nil
	}
}

func (db *ConfirmationsDocstoreDatabase) getConfirmationsWithCallback(ctx context.Context, q *docstore.Query, cb ListConfirmationsFunc) error {

	iter := q.Get(ctx)
	defer iter.Stop()

	for {

		var c confirmation.Confirmation
		err := iter.Next(ctx, &c)

		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("Failed to interate, %w", err)
		} else {

			err := cb(ctx, &c)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
