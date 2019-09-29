package mailinglist

import (
	"errors"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/database/fs"
	"github.com/aaronland/go-mailinglist/sender"
	"github.com/aaronland/go-string/dsn"
	"github.com/aaronland/gomail"
	_ "log"
	"strings"
)

func NewSenderFromDSN(str_dsn string) (gomail.Sender, error) {

	dsn_map, err := dsn.StringToDSNWithKeys(str_dsn, "sender")

	if err != nil {
		return nil, err
	}

	var s gomail.Sender

	switch strings.ToUpper(dsn_map["sender"]) {
	case "SMTP":
		s, err = sender.NewSMTPSenderFromDSN(str_dsn)
	case "STDOUT":
		s, err = sender.NewStdoutSender()
	default:
		err = errors.New("Invalid sender")
	}

	if err != nil {
		return nil, err
	}

	return s, nil

}

func NewSubscriptionsDatabaseFromDSN(str_dsn string) (database.SubscriptionsDatabase, error) {

	dsn_map, err := dsn.StringToDSNWithKeys(str_dsn, "database")

	if err != nil {
		return nil, err
	}

	var db database.SubscriptionsDatabase

	switch strings.ToUpper(dsn_map["database"]) {
	case "FS":

		root, ok := dsn_map["root"]

		if ok {
			db, err = fs.NewFSSubscriptionsDatabase(root)
		} else {
			err = errors.New("Missing 'root' DSN string")
		}

	default:
		err = errors.New("Invalid database")
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewInvitationsDatabaseFromDSN(str_dsn string) (database.InvitationsDatabase, error) {

	dsn_map, err := dsn.StringToDSNWithKeys(str_dsn, "database")

	if err != nil {
		return nil, err
	}

	var db database.InvitationsDatabase

	switch strings.ToUpper(dsn_map["database"]) {
	case "FS":

		root, ok := dsn_map["root"]

		if ok {
			db, err = fs.NewFSInvitationsDatabase(root)
		} else {
			err = errors.New("Missing 'root' DSN string")
		}

	default:
		err = errors.New("Invalid database")
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewConfirmationsDatabaseFromDSN(str_dsn string) (database.ConfirmationsDatabase, error) {

	dsn_map, err := dsn.StringToDSNWithKeys(str_dsn, "database")

	if err != nil {
		return nil, err
	}

	var db database.ConfirmationsDatabase

	switch strings.ToUpper(dsn_map["database"]) {
	case "FS":

		root, ok := dsn_map["root"]

		if ok {
			db, err = fs.NewFSConfirmationsDatabase(root)
		} else {
			err = errors.New("Missing 'root' DSN string")
		}

	default:
		err = errors.New("Invalid database")
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewEventLogsDatabaseFromDSN(str_dsn string) (database.EventLogsDatabase, error) {

	dsn_map, err := dsn.StringToDSNWithKeys(str_dsn, "database")

	if err != nil {
		return nil, err
	}

	var db database.EventLogsDatabase

	switch strings.ToUpper(dsn_map["database"]) {
	case "FS":

		root, ok := dsn_map["root"]

		if ok {
			db, err = fs.NewFSEventLogsDatabase(root)
		} else {
			err = errors.New("Missing 'root' DSN string")
		}

	default:
		err = errors.New("Invalid database")
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
