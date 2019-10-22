package http

import (
	"log"
)

// 01 - 99 (feature flags)

const DISABLED_SUBSCRIBE int = 1
const DISABLED_UNSUBSCRIBE int = 2
const DISABLED_CONFIRM int = 3
const DISABLED_INVITE int = 4

const DISABLED_SUBSCRIBE_MESSAGE string = "Subscriptions are not currently enabled."
const DISABLED_UNSUBSCRIBE_MESSAGE string = "Unsubscribing is not currently enabled."

var errors_map = map[int]string {
	-1:                   "Unknown error",
	DISABLED_SUBSCRIBE:   DISABLED_SUBSCRIBE_MESSAGE,
	DISABLED_UNSUBSCRIBE: DISABLED_UNSUBSCRIBE_MESSAGE,
}

type MailingListError struct {
	Code    int
	Message string
}

func (e *MailingListError) Error() string {
	return e.Message
}

func NewMailingListError(code int, err error) *MailingListError {

	msg, ok := errors_map[code]

	if !ok {
		code = -1
		msg, _ = errors_map[code]
	}

	log.Printf("[ERROR][%d] %v\n", code, err)

	return &MailingListError{
		Code:    code,
		Message: msg,
	}
}

// 01 - 99 (feature flags)
// subscription(s) disabled
// unsubscribe disabled
// invitation(s) disabled

// 100 - 199
// params sanitize error
// params missing param

// 200 - 299
// email invalid address
// email new message error
// email send error

// 300 - 399

// subscription already exists <-- how to not leak data?
// subscription retrieve error
// subscription not found
// subscription create (new) error
// subscription confirm error
// subscription disabled error
// subscription add (to db) error

// 400 - 499
// confirmation retrieve error
// confirmation expired error
// confirmation create error
// confirmation add (to db) error
// confirmation invalid action error

// 500 - 599
// invitation retrieve / not found
// invitation unavailable
// invitation accept error
// invitation update (db) error
// invitation unconfirmed <-- ??
// invitation list (for inviter)
// invitation max invites
// invitation new invite
// invitation add invite

// events add log error
