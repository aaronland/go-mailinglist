package fs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aaronland/go-mailinglist/confirmation"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/whosonfirst/walk"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type FSConfirmationsDatabase struct {
	database.ConfirmationsDatabase
	root string
}

func NewFSConfirmationsDatabase(root string) (database.ConfirmationsDatabase, error) {

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

	db := FSConfirmationsDatabase{
		root: abs_root,
	}

	return &db, nil
}

func (db *FSConfirmationsDatabase) AddConfirmation(conf *confirmation.Confirmation) error {

	path := db.pathForConfirmation(conf)

	_, err := os.Stat(path)

	if err == nil {
		return nil
	}

	return db.writeConfirmation(conf, path)
}

func (db *FSConfirmationsDatabase) RemoveConfirmation(conf *confirmation.Confirmation) error {

	path := db.pathForConfirmation(conf)

	_, err := os.Stat(path)

	if err != nil {

		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	return os.Remove(path)
}

func (db *FSConfirmationsDatabase) GetConfirmationWithAddress(code string) (*confirmation.Confirmation, error) {

	path := db.pathForCode(code)

	_, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	return db.readConfirmation(path)
}

func (db *FSConfirmationsDatabase) readConfirmation(path string) (*confirmation.Confirmation, error) {

	fh, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer fh.Close()

	body, err := ioutil.ReadAll(fh)

	if err != nil {
		return nil, err
	}

	var conf *confirmation.Confirmation

	err = json.Unmarshal(body, &conf)

	if err != nil {
		return nil, err
	}

	return conf, nil
}

func (db *FSConfirmationsDatabase) writeConfirmation(conf *confirmation.Confirmation, path string) error {

	enc, err := json.Marshal(conf)

	if err != nil {
		return err
	}

	fh, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0600)

	if err != nil {
		return err
	}

	fh.Write(enc)
	return fh.Close()
}

func (db *FSConfirmationsDatabase) crawlConfirmations(ctx context.Context, cb database.ListConfirmationsFunc) error {

	walker := func(path string, info os.FileInfo, err error) error {

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".json") {
			return nil
		}

		conf, err := db.readConfirmation(path)

		if err != nil {
			return err
		}

		return cb(conf)
	}

	return walk.Walk(db.root, walker)
}

func (db *FSConfirmationsDatabase) pathForConfirmation(conf *confirmation.Confirmation) string {
	return db.pathForCode(conf.Code)
}

func (db *FSConfirmationsDatabase) pathForCode(code string) string {
	fname := fmt.Sprintf("%s.json", code)
	return filepath.Join(db.root, fname)
}
