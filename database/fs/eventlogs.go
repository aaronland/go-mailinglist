package fs

import (
	"errors"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/eventlog"
	"os"
	"path/filepath"
)

type FSEventLogsDatabase struct {
	database.EventLogsDatabase
	root string
}

func NewFSEventLogsDatabase(root string) (database.EventLogsDatabase, error) {

	abs_root, err := filepath.Abs(root)

	if err != nil {
		return nil, err
	}

	info, err := os.Stat(abs_root)

	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, errors.New("Root is not a directory")
	}

	/*
		if info.Mode() != 0700 {
			return nil, errors.New("Root permissions must be 0700")
		}
	*/

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
