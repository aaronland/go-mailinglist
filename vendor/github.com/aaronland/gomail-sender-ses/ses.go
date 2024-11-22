package ses

// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/ses-example-send-email.html
// https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-email-addresses-procedure.html

// https://docs.aws.amazon.com/sdk-for-go/api/service/ses/#SES.VerifyEmailIdentity
// https://docs.aws.amazon.com/sdk-for-go/api/service/ses/#SES.WaitUntilIdentityExists
// https://docs.aws.amazon.com/sdk-for-go/api/service/ses/#SES.ListIdentities
// https://docs.aws.amazon.com/sdk-for-go/api/service/ses/#SES.DeleteIdentity

// https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-personalized-email-api.html
// https://docs.aws.amazon.com/ses/latest/APIReference/API_CreateCustomVerificationEmailTemplate.html
// https://docs.aws.amazon.com/sdk-for-go/api/service/ses/#SES.CreateCustomVerificationEmailTemplate

// https://docs.aws.amazon.com/ses/latest/DeveloperGuide/request-production-access.html
// https://us-west-2.console.aws.amazon.com/ses/home?region=us-west-2#smtp-settings:

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/url"
	"sync"

	"github.com/aaronland/go-aws-auth"
	"github.com/aaronland/gomail-sender"
	"github.com/aaronland/gomail/v2"
	aws_ses "github.com/aws/aws-sdk-go-v2/service/ses"
	aws_ses_types "github.com/aws/aws-sdk-go-v2/service/ses/types"
)

// SESSender implements the `gomail.Sender` inferface for delivery messages using the AWS Simple Email Service (SES).
type SESSender struct {
	gomail.Sender
	client *aws_ses.Client
}

// In principle this could also be done with a sync.OnceFunc call but that will
// require that everyone uses Go 1.21 (whose package import changes broke everything)
// which is literally days old as I write this. So maybe a few releases after 1.21

var register_mu = new(sync.RWMutex)
var register_map = map[string]bool{}

func init() {

	ctx := context.Background()
	err := RegisterSchemes(ctx)

	if err != nil {
		panic(err)
	}
}

// RegisterSchemes will explicitly register all the schemes associated with the `SESSender`.
func RegisterSchemes(ctx context.Context) error {

	roster := map[string]sender.SenderInitializeFunc{
		"ses": NewSESSender,
	}

	register_mu.Lock()
	defer register_mu.Unlock()

	for scheme, fn := range roster {

		_, exists := register_map[scheme]

		if exists {
			continue
		}

		err := sender.RegisterSender(ctx, scheme, fn)

		if err != nil {
			return fmt.Errorf("Failed to register sender for '%s', %w", scheme, err)
		}

		register_map[scheme] = true
	}

	return nil
}

// NewSESSender returns a new `SESSender` instance for delivering messages using the AWS Simple Email Service (SES),
// configured by 'uri' which is expected to take the form of:
//
//	ses://?credentials={CREDENTIALS}&region={REGION}
//
// Where: {CREDENTIALS} is a valid `aaronland/go-aws-session` credentials string; {REGION} is a valid AWS region.
func NewSESSender(ctx context.Context, uri string) (gomail.Sender, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URL, %w", err)
	}

	q := u.Query()

	credentials := q.Get("credentials")
	region := q.Get("region")

	cfg_uri := fmt.Sprintf("aws://%s?credentials=%s", region, credentials)

	cfg, err := auth.NewConfig(ctx, cfg_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create new AWS config, %w", err)
	}

	cl := aws_ses.NewFromConfig(cfg)

	s := SESSender{
		client: cl,
	}

	// https://docs.aws.amazon.com/sdk-for-go/api/service/ses/#GetSendQuotaOutput

	return &s, nil
}

// Send will deliver 'msg' to each recipient listed in 'to' using the AWS Simple Email Service (SES).
func (s *SESSender) Send(from string, to []string, msg io.WriterTo) error {

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	_, err := msg.WriteTo(wr)

	if err != nil {
		return fmt.Errorf("Failed to write message to buffer, %w", err)
	}

	wr.Flush()

	raw_msg := &aws_ses_types.RawMessage{
		Data: buf.Bytes(),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, recipient := range to {

		err := s.sendMessage(ctx, from, recipient, raw_msg)

		// maybe check err here and sometimes continue ?

		if err != nil {
			return fmt.Errorf("Failed to send message, %w", err)
		}
	}

	return nil
}

// Send will deliver 'msg' to 'recipient' using the AWS Simple Email Service (SES).
func (s *SESSender) sendMessage(ctx context.Context, sender string, recipient string, msg *aws_ses_types.RawMessage) error {

	// throttle send here... (see quota stuff above)

	select {
	case <-ctx.Done():
		return nil
	default:
		// pass
	}

	req := &aws_ses.SendRawEmailInput{
		RawMessage: msg,
	}

	_, err := s.client.SendRawEmail(ctx, req)

	if err != nil {
		return fmt.Errorf("Failed to send message with SES, %w", err)
	}

	return nil
}
