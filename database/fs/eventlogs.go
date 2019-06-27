package fs

import (
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

		if !os.IsNotExist(err){
			return err
		}

		err = os.MkdirAll(root, 0700)

		if err != nil {
			return err
		}
	}
	
	fname := fmt.Sprintf("%d-%s.json", ev.Created, ev.Event)
	path := filepath.Join(root, fname)

	_, err = os.Stat(path)

	if err == nil {
		return nil
	}

	return marshalData(ev, path)
}
