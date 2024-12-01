# gomail-sender-ses

Go package to implement the `gomail.Sender` interface using the AWS Simple Email Service (SES).

## Documentation

Documentation is incomplete.

[![Go Reference](https://pkg.go.dev/badge/github.com/aaronland/gomail-sender-ses.svg)](https://pkg.go.dev/github.com/aaronland/gomail-sender-ses)

## Example

```
package main

import(
	"context"
	"fmt"
	"net/url"

	"github.com/aaronland/gomail-sender"
	_ "github.com/aaronland/gomail-sender-ses/v2"	
	"github.com/aaronland/gomail/v2"	
)

func main() {

	sender_uri := flag.String("sender-uri", "", "A valid aaronland/gomail-sender URI")
	from := flag.String("from", "", "A valid From: address (that has been registered with SES)")
	to := flag.String("to", "", "A valid To: address")
	subject := flag.String("subject", "", "A valid email subject")		
	
	flag.Parse()
	
	ctx := context.Background()

	mail_sender, _ := sender.NewSender(ctx, *sender_uri)

	msg := gomail.NewMessage()

	msg.SetHeader("Subject", *subject)
	msg.SetHeader("From", *from)
	msg.SetHeader("To", *to)

	msg.SetBody("text/plain", "This message left intentionally blank.")

	gomail.Send(mail_sender, msg)
}
```

_Error handling removed for the sake of brevity._

## Tools

```
$> make cli
go build -mod vendor -o bin/send cmd/send/main.go
```

### send

```
$> ./bin/send \
	-sender-uri 'ses://?credentials={CREDENTIALS}&region={REGION}'
	-from bob@example.com \
	-to alice@example.com \
	-subject 'This is a test'
```

## Sender URIs

Sender URIs take the form of:

```
ses://?credentials={CREDENTIALS}&region={REGION}
```

Where `{CREDENTIALS}` is a valid `aaronland/go-aws-session` credentials string and `{REGION}` is a valid AWS region.

Credentials strings for AWS sessions are defined as string labels. They are:

| Label | Description |
| --- | --- |
| `env:` | Read credentials from AWS defined environment variables. |
| `iam:` | Assume AWS IAM credentials are in effect. |
| `{AWS_PROFILE_NAME}` | Use the profile from the default AWS credentials location. |
| `{AWS_CREDENTIALS_PATH}:{AWS_PROFILE_NAME}` | Use the profile from a user-defined AWS credentials location. |

## See also

* https://github.com/aaronland/gomail/v2
* https://github.com/aaronland/gomail-sender
* https://docs.aws.amazon.com/sdk-for-go/api/service/ses
* https://github.com/aaronland/go-aws-session
=======
     	config_uri := "aws://?region={REGION}&credentials={CREDENTIALS}"
	sender_uri := fmt.Sprintf("ses://?config-uri=%s", url.QueryEscape)

	from := "bob@bob.com"
	to := "alice@alice.com"

	ctx := context.Background()
	
	s, _ := sender.NewSender(ctx, sender_uri)

	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)	
	msg.SetHeader("This is a test")
	msg.SetBody("text/plain", "This is a test")

	gomail.Send(s, msg)
```

## Credentials

Credentials strings (as `?credentials={CREDENTIALS}`) are expected to be a valid [aaronland/go-aws-auth](https://github.com/aaronland/go-aws-auth?tab=readme-ov-file#credentials) credentials "label":

| Label | Description |
| --- | --- |
| `anon:` | Empty or anonymous credentials. |
| `env:` | Read credentials from AWS defined environment variables. |
| `iam:` | Assume AWS IAM credentials are in effect. |
| `sts:{ARN}` | Assume the role defined by `{ARN}` using STS credentials. |
| `{AWS_PROFILE_NAME}` | This this profile from the default AWS credentials location. |
| `{AWS_CREDENTIALS_PATH}:{AWS_PROFILE_NAME}` | This this profile from a user-defined AWS credentials location. |

For example:

```
aws:///us-east-1?credentials=iam:
```

## See also

* https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sesv2
* https://github.com/aaronland/gomail-sender
* https://github.com/aaronland/gomail/v2
* https://github.com/aaronland/go-aws-auth
