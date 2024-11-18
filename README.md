># go-mailinglist

This is work-in-progress.

## Example

```
package main

import (
	"net/mail"

	"github.com/aaronland/go-mailinglist/v2/message"	
	"github.com/aaronland/gomail-sender"
	"github.com/aaronland/gomail/v2"	
)

func main() {

	ctx := context.Background()
	s, _ := sender.NewSender(ctx, "stdout://")

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

	message.SendMessage(ctx, m, opts)
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

## See also

* https://github.com/aaronland/gomail
* https://github.com/aaronland/gomail-ses
