# gomail-sender

Package sender provides a common interface for creating new `aaronland/gomail/v2` instances using a URI-based syntax.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/aaronland/gomail-sender.svg)](https://pkg.go.dev/github.com/aaronland/gomail-sender)

## Example

```
import (
	"context"
	"github.com/aaronland/gomail/v2"		
)

func main(){

	ctx := context.Background()
	s, _ := NewSender(ctx, "stdout://")

	msg := gomail.NewMessage()
	msg.SetBody("text/plain", "Hello world.")
	msg.SetHeader("From", "from@example.com")
	msg.SetHeader("To", "to@example.com")
	msg.SetHeader("Subject", "Stdout sender")
	
	gomail.Send(s, msg)
}
```

_Error handling removed for the sake of brevity._

## See also

* https://github.com/aaronland/gomail