package ses

// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sesv2

import (
	"bufio"
	"bytes"
	"context"
	_ "fmt"
	"io"
	"net/url"

	"github.com/aaronland/go-aws-auth"
	"github.com/aaronland/gomail-sender"
	"github.com/aaronland/gomail/v2"
	"github.com/aws/aws-sdk-go-v2/aws"
	aws_ses "github.com/aws/aws-sdk-go-v2/service/sesv2"
	aws_ses_types "github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

type SESSender struct {
	gomail.Sender
	client *aws_ses.Client
}

func init() {
	ctx := context.Background()
	err := sender.RegisterSender(ctx, "ses", NewSESSender)

	if err != nil {
		panic(err)
	}
}

func NewSESSender(ctx context.Context, uri string) (gomail.Sender, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	q := u.Query()

	config_uri := q.Get("config-uri")

	cfg, err := auth.NewConfig(ctx, config_uri)

	if err != nil {
		return nil, err
	}

	cl := aws_ses.NewFromConfig(cfg)

	s := SESSender{
		client: cl,
	}

	return &s, nil
}

func (s *SESSender) Send(from string, to []string, msg io.WriterTo) error {

	ctx := context.Background()

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	_, err := msg.WriteTo(wr)

	if err != nil {
		return err
	}

	wr.Flush()

	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sesv2#Client.SendEmail
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sesv2#SendEmailInput
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sesv2@v1.38.3/types#EmailContent
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sesv2@v1.38.3/types#RawMessage

	raw_msg := &aws_ses_types.RawMessage{
		Data: buf.Bytes(),
	}

	content := &aws_ses_types.EmailContent{
		Raw: raw_msg,
	}

	dest := &aws_ses_types.Destination{
		ToAddresses: to,
	}

	req := &aws_ses.SendEmailInput{
		Content:          content,
		Destination:      dest,
		FromEmailAddress: aws.String(from),
	}

	_, err = s.client.SendEmail(ctx, req)

	if err != nil {
		return err
	}

	return nil
}
