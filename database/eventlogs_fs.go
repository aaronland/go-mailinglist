package database

import (
	"context"
	"fmt"
	"github.com/aaronland/go-mailinglist/eventlog"
	_ "log"
	"net/url"
	"os"
	"path/filepath"
)

type FSEventLogsDatabase struct {
	EventLogsDatabase
	root string
}

func init() {
	ctx := context.Background()
	RegisterEventLogsDatabase(ctx, "fs", NewFSEventLogsDatabase)
}

func NewFSEventLogsDatabase(ctx context.Context, uri string) (EventLogsDatabase, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	root := u.Path

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

func (db *FSEventLogsDatabase) ListEventLogs(ctx context.Context, callback ListEventLogsFunc) error {

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
