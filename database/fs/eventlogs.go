package fs

import (
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/eventlog"
	"os"
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

func (db *FSEventLogsDatabase) AddEventLog(sub *eventlog.EventLog) error {

	path := "fix me"

	_, err := os.Stat(path)

	if err == nil {
		return nil
	}

	return marshalData(sub, path)
}
