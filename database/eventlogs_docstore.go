package database

import (
	"context"
	"fmt"
	"io"

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
	q := db.collection.Query()
	return db.getEventLogsWithCallback(ctx, q, cb)
}

func (db *EventLogsDocstoreDatabase) Close() error {
	return db.collection.Close()
}

func (db *EventLogsDocstoreDatabase) getDeliveryWithQuery(ctx context.Context, q *docstore.Query) (*eventlog.EventLog, error) {

	iter := q.Get(ctx)
	defer iter.Stop()

	var l eventlog.EventLog
	err := iter.Next(ctx, &l)

	if err == io.EOF {
		return nil, NoRecordError("")
	} else if err != nil {
		return nil, fmt.Errorf("Failed to interate, %w", err)
	} else {
		return &l, nil
	}
}

func (db *EventLogsDocstoreDatabase) getEventLogsWithCallback(ctx context.Context, q *docstore.Query, cb ListEventLogsFunc) error {

	iter := q.Get(ctx)
	defer iter.Stop()

	for {

		var l eventlog.EventLog
		err := iter.Next(ctx, &l)

		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("Failed to interate, %w", err)
		} else {

			err := cb(ctx, &l)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
