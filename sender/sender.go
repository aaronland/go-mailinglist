package sender

import (
	"errors"
	"github.com/aaronland/go-string/dsn"
	"github.com/aaronland/gomail"
	"github.com/aaronland/gomail-ses"
	"strings"
)

func NewSenderFromDSN(str_dsn string) (gomail.Sender, error) {

	dsn_map, err := dsn.StringToDSNWithKeys(str_dsn, "sender")

	if err != nil {
		return nil, err
	}

	var sender gomail.Sender

	switch strings.ToUpper(dsn_map["sender"]) {
	case "SES":
		sender, err = ses.NewSESSender(str_dsn)
	default:
		err = errors.New("Invalid sender")
	}

	if err != nil {
		return nil, err
	}

	return sender, nil
}
