package sender

import (
	"context"
	"github.com/aaronland/gomail/v2"
	"net/url"
	"strconv"
)

func init() {

	ctx := context.Background()
	err := RegisterSender(ctx, "smtp", NewSMTPSender)

	if err != nil {
		panic(err)
	}
}

func NewSMTPSender(ctx context.Context, uri string) (gomail.Sender, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	q := u.Query()

	host := q.Get("host")
	str_port := q.Get("port")
	username := q.Get("username")
	password := q.Get("password")

	port, err := strconv.Atoi(str_port)

	if err != nil {
		return nil, err

	}

	d := gomail.NewDialer(host, port, username, password)
	return d.Dial()
}
