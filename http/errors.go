package http

import (
	"fmt"
	"log"
)

// -1 (house keeping)

const E_UNKNOWN_ERROR int = -1
const E_UNKNOWN_ERROR_MESSAGE = "Unknown error"

// 01 - 99 (feature flags)

const E_DISABLED_SUBSCRIBE int = 1
const E_DISABLED_UNSUBSCRIBE int = 2
const E_DISABLED_CONFIRM int = 3
const E_DISABLED_INVITE int = 4

const E_DISABLED_SUBSCRIBE_MESSAGE string = "Subscriptions are not currently enabled."
const E_DISABLED_UNSUBSCRIBE_MESSAGE string = "Unsubscribing is not currently enabled."
const E_DISABLED_CONFIRM_MESSAGE string = "Confirmations are not currently enabled."
const E_DISABLED_INVITE_MESSAGE string = "Invites are not currently enabled."

// 100 - 199 (input validation)

const E_INPUT_MISSING int = 100
const E_INPUT_PARSE int = 101

const E_INPUT_MISSING_MESSAGE string = "Missing parameter '%s'"
const E_INPUT_PARSE_MESSAGE string = "Invalid parameter '%s'"

// 200 - 299 (email)

const E_EMAIL_INVALID int = 200
const E_EMAIL_CREATE int = 201
const E_EMAIL_SEND int = 202

const E_EMAIL_INVALID_MESSAGE string = "Invalid email address '%s'."
const E_EMAIL_CREATE_MESSAGE string = "Failed to create new email message."
const E_EMAIL_SEND_MESSAGE string = "Failed to send email message."

// 300 - 399 (subscriptions)

const E_SUBSCRIPTION_EXISTS int = 300
const E_SUBSCRIPTION_RETRIEVE int = 301
const E_SUBSCRIPTION_NOTFOUND int = 302
const E_SUBSCRIPTION_CREATE int = 303
const E_SUBSCRIPTION_CONFIRM int = 304
const E_SUBSCRIPTION_DISABLED int = 305
const E_SUBSCRIPTION_ADD int = 306

const E_SUBSCRIPTION_EXISTS_MESSAGE string = "Already subscribed." // how to not leak data...?
const E_SUBSCRIPTION_RETRIEVE_MESSAGE string = "Failed to retrieve subscription."
const E_SUBSCRIPTION_NOTFOUND_MESSAGE string = "Subscription not found."
const E_SUBSCRIPTION_CREATE_MESSAGE string = "Failed to create new subscription."
const E_SUBSCRIPTION_CONFIRM_MESSAGE string = "Failed to confirm subscription."
const E_SUBSCRIPTION_DISABLED_MESSAGE string = "Subscription is disabled."
const E_SUBSCRIPTION_ADD_MESSAGE string = "Failed to add subscription."

// 400 - 499 (confirmations)

const E_CONFIRMATION_RETRIEVE int = 400
const E_CONFIRMATION_EXPIRED int = 401
const E_CONFIRMATION_CREATE int = 402
const E_CONFIRMATION_ADD int = 403
const E_CONFIRMATION_INVALID int = 404

const E_CONFIRMATION_RETRIEVE_MESSAGE string = "Failed to retrieve confirmation."
const E_CONFIRMATION_EXPIRED_MESSAGE string = "Confirmation has expired."
const E_CONFIRMATION_CREATE_MESSAGE string = "Failed to create confirmation."
const E_CONFIRMATION_ADD_MESSAGE string = "Failed to add confirmation."
const E_CONFIRMATION_INVALID_MESSAGE string = "Invalid confirmation."

var errors_map = map[int]string{
	E_UNKNOWN_ERROR:         E_UNKNOWN_ERROR_MESSAGE,
	E_DISABLED_SUBSCRIBE:    E_DISABLED_SUBSCRIBE_MESSAGE,
	E_DISABLED_UNSUBSCRIBE:  E_DISABLED_UNSUBSCRIBE_MESSAGE,
	E_DISABLED_CONFIRM:      E_DISABLED_CONFIRM_MESSAGE,
	E_DISABLED_INVITE:       E_DISABLED_INVITE_MESSAGE,
	E_INPUT_MISSING:         E_INPUT_MISSING_MESSAGE,
	E_INPUT_PARSE:           E_INPUT_PARSE_MESSAGE,
	E_EMAIL_INVALID:         E_EMAIL_INVALID_MESSAGE,
	E_EMAIL_CREATE:          E_EMAIL_CREATE_MESSAGE,
	E_EMAIL_SEND:            E_EMAIL_SEND_MESSAGE,
	E_SUBSCRIPTION_EXISTS:   E_SUBSCRIPTION_EXISTS_MESSAGE,
	E_SUBSCRIPTION_RETRIEVE: E_SUBSCRIPTION_RETRIEVE_MESSAGE,
	E_SUBSCRIPTION_NOTFOUND: E_SUBSCRIPTION_NOTFOUND_MESSAGE,
	E_SUBSCRIPTION_CREATE:   E_SUBSCRIPTION_CREATE_MESSAGE,
	E_SUBSCRIPTION_CONFIRM:  E_SUBSCRIPTION_CONFIRM_MESSAGE,
	E_SUBSCRIPTION_DISABLED: E_SUBSCRIPTION_DISABLED_MESSAGE,
	E_SUBSCRIPTION_ADD:      E_SUBSCRIPTION_ADD_MESSAGE,
	E_CONFIRMATION_RETRIEVE: E_CONFIRMATION_RETRIEVE_MESSAGE,
	E_CONFIRMATION_EXPIRED:  E_CONFIRMATION_EXPIRED_MESSAGE,
	E_CONFIRMATION_CREATE:   E_CONFIRMATION_CREATE_MESSAGE,
	E_CONFIRMATION_ADD:      E_CONFIRMATION_ADD_MESSAGE,
	E_CONFIRMATION_INVALID:  E_CONFIRMATION_INVALID_MESSAGE,
}

type ApplicationError struct {
	Code    int
	Message string
	Detail  string
}

func (e *ApplicationError) Error() string {

	if e.Detail == "" {
		return e.Message
	}

	return fmt.Sprintf("%s: %s", e.Message, e.Detail)
}

func NewFeatureFlagError(code int) *ApplicationError {
	return NewApplicationError(code, "", nil)
}

func NewParamsError(code int, param string) *ApplicationError {
	return NewApplicationError(code, param, nil)
}

func NewApplicationError(code int, detail string, err error) *ApplicationError {

	msg, ok := errors_map[code]

	if !ok {
		code = -1
		msg, _ = errors_map[code]
	}

	log.Printf("[ERROR][%d] %s (%v)\n", code, detail, err)

	return &ApplicationError{
		Code:    code,
		Message: msg,
		Detail:  detail,
	}
}

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
