package fs

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/aaronland/go-mailinglist/confirmation"
	"github.com/aaronland/go-mailinglist/eventlog"
	"github.com/aaronland/go-mailinglist/subscription"
	"github.com/whosonfirst/walk"
	"io/ioutil"
	"os"
	_ "path/filepath"
	"strings"
)

func marshalData(data interface{}, path string) error {

	enc, err := json.Marshal(data)

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

func unmarshalData(path string, data_type string) (interface{}, error) {

	switch data_type {
	case "confirmations", "eventlogs", "subscriptions":
		// pass
	default:
		return nil, errors.New("Unsupported interface")
	}

	fh, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer fh.Close()

	body, err := ioutil.ReadAll(fh)

	if err != nil {
		return nil, err
	}

	var data interface{}

	switch data_type {

	case "confirmations":

		var conf *confirmation.Confirmation
		err = json.Unmarshal(body, &conf)

		if err == nil {
			data = conf
		}

	case "eventlogs":

		var log *eventlog.EventLog
		err = json.Unmarshal(body, &log)

		if err == nil {
			data = log
		}

	case "subscriptions":

		var sub *subscription.Subscription
		err = json.Unmarshal(body, &sub)

		if err == nil {
			data = sub
		}

	default:
		err = errors.New("Unsupported data type")
	}

	return data, err
}

func crawlDatabase(ctx context.Context, root string, cb func(context.Context, string) error) error {

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

		return cb(ctx, path)
	}

	return walk.Walk(root, walker)
}
