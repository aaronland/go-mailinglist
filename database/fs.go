package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aaronland/go-mailinglist/confirmation"	
	"github.com/aaronland/go-mailinglist/subscription"
	"github.com/whosonfirst/walk"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type FSSubscriptionDatabase struct {
	SubscriptionDatabase
	root string
}

type FSConfirmationDatabase struct {
	ConfirmationDatabase
	root string
}

func NewFSSubscriptionDatabase(root string) (SubscriptionDatabase, error) {

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

	db := FSSubscriptionDatabase{
		root: abs_root,
	}

	return &db, nil
}

func (db *FSSubscriptionDatabase) AddSubscription(sub *subscription.Subscription) error {

	path := db.pathForSubscription(sub)

	_, err := os.Stat(path)

	if err == nil {
		return nil
	}

	return db.writeSubscription(sub, path)
}

func (db *FSSubscriptionDatabase) RemoveSubscription(sub *subscription.Subscription) error {

	path := db.pathForSubscription(sub)

	_, err := os.Stat(path)

	if err != nil {

		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	return os.Remove(path)
}

func (db *FSSubscriptionDatabase) UpdateSubscription(sub *subscription.Subscription) error {

	path := db.pathForSubscription(sub)

	_, err := os.Stat(path)

	if err == nil {
		return nil
	}

	return db.writeSubscription(sub, path)
}

func (db *FSSubscriptionDatabase) GetSubscriptionWithAddress(addr string) (*subscription.Subscription, error) {

	path := db.pathForAddress(addr)

	_, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	return db.readSubscription(path)
}

func (db *FSSubscriptionDatabase) readSubscription(path string) (*subscription.Subscription, error) {

	fh, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer fh.Close()

	body, err := ioutil.ReadAll(fh)

	if err != nil {
		return nil, err
	}

	var sub *subscription.Subscription

	err = json.Unmarshal(body, &sub)

	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (db *FSSubscriptionDatabase) writeSubscription(sub *subscription.Subscription, path string) error {

	enc, err := json.Marshal(sub)

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

func (db *FSSubscriptionDatabase) ConfirmedSubscriptions(ctx context.Context, cb func(*subscription.Subscription) error) error {

	local_cb := func(ctx context.Context, sub *subscription.Subscription) error {

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		if !sub.Confirmed {
			return nil
		}

		return cb(sub)
	}

	return db.crawlSubscriptions(ctx, local_cb)
}

func (db *FSSubscriptionDatabase) UnconfirmedSubscriptions(ctx context.Context, cb func(*subscription.Subscription) error) error {

	local_cb := func(ctx context.Context, sub *subscription.Subscription) error {

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		if sub.Confirmed {
			return nil
		}

		return cb(sub)
	}

	return db.crawlSubscriptions(ctx, local_cb)
}

func (db *FSSubscriptionDatabase) crawlSubscriptions(ctx context.Context, cb func(context.Context, *subscription.Subscription) error) error {

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

		sub, err := db.readSubscription(path)

		if err != nil {
			return err
		}

		return cb(ctx, sub)
	}

	return walk.Walk(db.root, walker)
}

func (db *FSSubscriptionDatabase) pathForSubscription(sub *subscription.Subscription) string {
	return db.pathForAddress(sub.Address)
}

func (db *FSSubscriptionDatabase) pathForAddress(addr string) string {
	fname := fmt.Sprintf("%s.json", addr)
	return filepath.Join(db.root, fname)
}

// confirmation

func NewFSConfirmationDatabase(root string) (ConfirmationDatabase, error) {

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

	db := FSConfirmationDatabase{
		root: abs_root,
	}

	return &db, nil
}

func (db *FSConfirmationDatabase) AddConfirmation(conf *confirmation.Confirmation) error {

	path := db.pathForConfirmation(conf)

	_, err := os.Stat(path)

	if err == nil {
		return nil
	}

	return db.writeConfirmation(sub, path)
}

func (db *FSConfirmationDatabase) RemoveConfirmation(conf *confirmation.Confirmation) error {

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

func (db *FSConfirmationDatabase) GetConfirmationWithAddress(code string) (*confirmation.Confirmation, error) {

	path := db.pathForCode(code)

	_, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	return db.readConfirmation(path)
}

func (db *FSConfirmationDatabase) readConfirmation(path string) (*confirmation.Confirmation, error) {

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

	err = json.Unmarshal(body, &sub)

	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (db *FSConfirmationDatabase) writeConfirmation(conf *confirmation.Confirmation, path string) error {

	enc, err := json.Marshal(sub)

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

func (db *FSConfirmationDatabase) crawlConfirmations(ctx context.Context, cb func(context.Context, *confirmation.Confirmation) error) error {

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

		sub, err := db.readConfirmation(path)

		if err != nil {
			return err
		}

		return cb(ctx, sub)
	}

	return walk.Walk(db.root, walker)
}

func (db *FSConfirmationDatabase) pathForConfirmation(sub *confirmation.Confirmation) string {
	return db.pathForCode(sub.Code)
}

func (db *FSConfirmationDatabase) pathForCode(code string) string {
	fname := fmt.Sprintf("%s.json", code)
	return filepath.Join(db.root, fname)
}
