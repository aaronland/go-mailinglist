package mailinglist

import (
	"github.com/aaronland/gomail"
)

type Mailer interface {
	SendMessage(m *gomail.Message, s *Subscriber) error
}
