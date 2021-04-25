package database

import (
	"context"
	"encoding/json"
	"github.com/aaronland/go-mailinglist/confirmation"
	"io/ioutil"
	"net/url"
	"os"
)

type FSConfirmationsDatabase struct {
	ConfirmationsDatabase
	root string
}

func NewFSConfirmationsDatabase(ctx context.Context, uri string) (ConfirmationsDatabase, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	root := u.Path

	abs_root, err := ensureRoot(root)

	if err != nil {
		return nil, err
	}

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

func (db *FSConfirmationsDatabase) GetConfirmationWithCode(code string) (*confirmation.Confirmation, error) {

	path := pathForAddress(db.root, code)

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

func (db *FSConfirmationsDatabase) crawlConfirmations(ctx context.Context, cb ListConfirmationsFunc) error {

	local_cb := func(ctx context.Context, path string) error {

		conf, err := db.readConfirmation(path)

		if err != nil {
			return err
		}

		return cb(conf)
	}

	return crawlDatabase(ctx, db.root, local_cb)
}

func (db *FSConfirmationsDatabase) pathForConfirmation(conf *confirmation.Confirmation) string {
	return pathForAddress(db.root, conf.Code)
}
