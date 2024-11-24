package invitation

import (
	"net/mail"
	"time"

	"github.com/aaronland/go-mailinglist/v2/code"
	"github.com/aaronland/go-mailinglist/v2/subscription"
)

const INVITATION_STATUS_AVAILABLE int = 0
const INVITATION_STATUS_ACCEPTED int = 1
const INVITATION_STATUS_DISABLED int = 2

type Invitation struct {
	Code         string `json:"code"`
	Inviter      string `json:"inviter"`
	Invitee      string `json:"invitee"`
	Created      int64  `json:"created"`
	LastModified int64  `json:"lastmodified"`
	Status       int    `json:"status"`
}

func NewInvitation(sub *subscription.Subscription) (*Invitation, error) {

	now := time.Now()
	ts := now.Unix()

	code, err := code.NewSecretCodeWithTime(now)

	if err != nil {
		return nil, err
	}

	invite := &Invitation{
		Code:         code,
		Inviter:      sub.Address,
		Invitee:      "",
		Created:      ts,
		LastModified: ts,
		Status:       INVITATION_STATUS_AVAILABLE,
	}

	return invite, nil
}

func (i *Invitation) IsAvailable() bool {

	if i.Status == INVITATION_STATUS_AVAILABLE {
		return true
	}

	return false
}

func (i *Invitation) Accept(invitee string) error {

	addr, err := mail.ParseAddress(invitee)

	if err != nil {
		return err
	}

	now := time.Now()
	ts := now.Unix()

	i.Invitee = addr.Address
	i.LastModified = ts
	i.Status = INVITATION_STATUS_ACCEPTED

	return nil
}
