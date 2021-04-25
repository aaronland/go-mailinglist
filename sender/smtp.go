package sender

import (
	"github.com/aaronland/go-string/dsn"
	"github.com/aaronland/gomail/v2"
	"strconv"
)

func NewSMTPSenderFromDSN(str_dsn string) (gomail.Sender, error) {

	dsn_map, err := dsn.StringToDSNWithKeys(str_dsn, "host", "port", "username", "password")

	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(dsn_map["port"])

	if err != nil {
		return nil, err
	}

	host := dsn_map["host"]
	username := dsn_map["username"]
	password := dsn_map["password"]

	return NewSMTPSender(host, port, username, password)
}

func NewSMTPSender(host string, port int, username string, password string) (gomail.Sender, error) {

	d := gomail.NewDialer(host, port, username, password)
	return d.Dial()
}
