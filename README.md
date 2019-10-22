># go-mailinglist

This is work-in-progress.

## Example

```
package main

import (
	"github.com/aaronland/go-mailinglist/message"	
	"github.com/aaronland/go-mailinglist/sender"
	"github.com/aaronland/gomail"	
	"net/mail"
)

func main() {

	s, _ := sender.NewStdoutSender()

	to, _ := mail.ParseAddress("to@example.com")
	from, _ := mail.ParseAddress("from@example.com")	
	subject := "This is the subject"

	opts := &message.SendMessageOptions{
		Sender: s,
		Subject: subject,
		From: from,
		To: to,
	}
	
	m := gomail.NewMessage()
	m.SetBody("text/html", "<p>hello world</p>")

	message.SendMessage(m, opts)
}
```

_Error handling removed for  brevity._

Which would produce this:

```
Mime-Version: 1.0
Date: Tue, 18 Jun 2019 15:59:04 -0700
From: from@example.com
To: to@example.com
Subject: This is the subject
Content-Type: text/html; charset=UTF-8
Content-Transfer-Encoding: quoted-printable

<p>hello world</p>
```

## Templates

### common_footer
### common_header

### confirm
### confirm_action
### confirm_email
### confirm_success

### crumb_error
### index

### invite_accept
### invite_code
### invite_how_it_works
### invite_request_email
### invite_request_success
### invite_request
### subscribe
### subscribe_success
### unsubscribe
### unsubscribe_success

## See also

* https://github.com/aaronland/gomail
* https://github.com/aaronland/gomail-ses
