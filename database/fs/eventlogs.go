package fs

import (
	"context"
	"fmt"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/eventlog"
	_ "log"
	"os"
	"path/filepath"
)

type FSEventLogsDatabase struct {
	database.EventLogsDatabase
	root string
}

func NewFSEventLogsDatabase(root string) (database.EventLogsDatabase, error) {

	abs_root, err := ensureRoot(root)

	if err != nil {
		return nil, err
	}

	db := FSEventLogsDatabase{
		root: abs_root,
	}

	return &db, nil
}

func (db *FSEventLogsDatabase) AddEventLog(ev *eventlog.EventLog) error {

	root := filepath.Join(db.root, ev.Address)

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

	fname := fmt.Sprintf("%d-%d.json", ev.Created, ev.Event)
	path := filepath.Join(root, fname)

	_, err = os.Stat(path)

	if err == nil {
		return nil
	}

	return marshalData(ev, path)
}

func (db *FSEventLogsDatabase) ListEventLogs(ctx context.Context, callback database.ListEventLogsFunc) error {

	local_cb := func(ctx context.Context, ev *eventlog.EventLog) error {

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		return callback(ev)
	}

	return db.crawlEventLogs(ctx, local_cb)
}

func (db *FSEventLogsDatabase) crawlEventLogs(ctx context.Context, cb func(ctx context.Context, ev *eventlog.EventLog) error) error {

	local_cb := func(ctx context.Context, path string) error {

		sub, err := db.readEventLog(path)

		if err != nil {
			return err
		}

		return cb(ctx, sub)
	}

	return crawlDatabase(ctx, db.root, local_cb)
}

func (db *FSEventLogsDatabase) readEventLog(path string) (*eventlog.EventLog, error) {

	ev, err := unmarshalData(path, "eventlog")

	if err != nil {
		return nil, err
	}

	return ev.(*eventlog.EventLog), nil
}
