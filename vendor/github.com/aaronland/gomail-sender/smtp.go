package sender

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/aaronland/gomail/v2"
)

func init() {

	ctx := context.Background()
	err := RegisterSender(ctx, "smtp", NewSMTPSender)

	if err != nil {
		panic(err)
	}
}

// NewSMTPSender returns a new `gomail.Sender` instance for delivery mail to a SMTP endpoint.
// 'uri' is expecteed to take the form of:
//
//	smtp://{HOST}?port={PORT}&username={USERNAME}&password={PASSWORD}
//
// Where ${HOST} is the name of the SMTP server host; {PORT} is the name of the SMTP server port; {USERNAME}
// and {PASSWORD} are the authentication credentials for accessing the SMTP server.
func NewSMTPSender(ctx context.Context, uri string) (gomail.Sender, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()

	host := q.Get("host")
	str_port := q.Get("port")
	username := q.Get("username")
	password := q.Get("password")

	port, err := strconv.Atoi(str_port)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse port number, %w", err)
	}

	d := gomail.NewDialer(host, port, username, password)
	return d.Dial()
}
