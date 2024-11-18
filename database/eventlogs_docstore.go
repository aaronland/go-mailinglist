package database

import (
	"context"

	"github.com/aaronland/go-mailinglist/v2/eventlog"
	aa_docstore "github.com/aaronland/gocloud-docstore"
	"gocloud.dev/docstore"
)

type EventLogsDocstoreDatabase struct {
	EventLogsDatabase
	collection *docstore.Collection
}

func init() {

	ctx := context.Background()

	err := RegisterEventLogsDatabase(ctx, "awsdynamodb", NewEventLogsDocstoreDatabase)

	if err != nil {
		panic(err)
	}

	for _, scheme := range docstore.DefaultURLMux().CollectionSchemes() {

		err := RegisterEventLogsDatabase(ctx, scheme, NewEventLogsDocstoreDatabase)

		if err != nil {
			panic(err)
		}

	}

}

func NewEventLogsDocstoreDatabase(ctx context.Context, uri string) (EventLogsDatabase, error) {

	col, err := aa_docstore.OpenCollection(ctx, uri)

	if err != nil {
		return nil, err
	}

	db := &EventLogsDocstoreDatabase{
		collection: col,
	}

	return db, nil
}

func (db *EventLogsDocstoreDatabase) AddEventLog(ctx context.Context, l *eventlog.EventLog) error {
	return db.collection.Put(ctx, l)
}

func (db *EventLogsDocstoreDatabase) ListEventLogs(ctx context.Context, cb ListEventLogsFunc) error {

}

func (db *EventLogsDocstoreDatabase) Close() error {
	return db.collection.Close()
}
