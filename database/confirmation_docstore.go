package database

import (
	"context"

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
}

func (db *ConfirmationsDocstoreDatabase) ListConfirmations(ctx context.Context, cb ListConfirmationsFunc) error {

}

func (db *ConfirmationsDocstoreDatabase) Close() error {
	return db.collection.Close()
}
