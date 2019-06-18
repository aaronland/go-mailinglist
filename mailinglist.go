package mailinglist

import (
	"context"
	"errors"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/database/fs"
	"github.com/aaronland/go-mailinglist/subscription"
	"github.com/aaronland/go-string/dsn"
	"github.com/aaronland/gomail"
	"strings"
)

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

func SendMailToList(sender gomail.Sender, db database.SubscriptionsDatabase, msg *gomail.Message) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return SendMailToListWithContext(ctx, sender, db, msg)
}

func SendMailToListWithContext(ctx context.Context, sender gomail.Sender, db database.SubscriptionsDatabase, msg *gomail.Message) error {

	cb := func(sub *subscription.Subscription) error {
		msg.SetHeader("To", sub.Address)
		return gomail.Send(sender, msg)
	}

	return db.ListSubscriptionsConfirmed(ctx, cb)
}
