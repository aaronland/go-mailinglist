package database

import (
	"context"
	"github.com/aaronland/go-mailinglist/block"
	"github.com/aaronland/go-mailinglist/confirmation"
	"github.com/aaronland/go-mailinglist/delivery"
	"github.com/aaronland/go-mailinglist/eventlog"
	"github.com/aaronland/go-mailinglist/subscription"
)

type NoRecordError string

func (err NoRecordError) Error() string {
	return "Database record does not exist"
}

func IsNotExist(err error) bool {

	switch err.(type) {
	case *NoRecordError, NoRecordError:
		return true
	default:
		return false
	}
}

type ListSubscriptionsFunc func(*subscription.Subscription) error
type ListConfirmationsFunc func(*confirmation.Confirmation) error
type ListEventLogsFunc func(*eventlog.EventLog) error
type ListDeliveriesFunc func(*delivery.Delivery) error
type ListBlockRulesFunc func(*block.Rule) error

type SubscriptionsDatabase interface {
	AddSubscription(*subscription.Subscription) error
	RemoveSubscription(*subscription.Subscription) error
	UpdateSubscription(*subscription.Subscription) error
	GetSubscriptionWithAddress(string) (*subscription.Subscription, error)
	ListSubscriptions(context.Context, ListSubscriptionsFunc) error
	ListSubscriptionsWithStatus(context.Context, ListSubscriptionsFunc, ...int) error
}

type ConfirmationsDatabase interface {
	AddConfirmation(*confirmation.Confirmation) error
	RemoveConfirmation(*confirmation.Confirmation) error
	GetConfirmationWithCode(string) (*confirmation.Confirmation, error)
	ListConfirmations(context.Context, ListConfirmationsFunc) error
}

type EventLogsDatabase interface {
	AddEventLog(*eventlog.EventLog) error
	ListEventLogs(context.Context, ListEventLogsFunc) error
}

type DeliveriesDatabase interface {
	AddDelivery(*delivery.Delivery) error
	ListDeliveries(context.Context, ListDeliveriesFunc) error
	GetDeliveryWithAddressAndMessageId(string, string) (*delivery.Delivery, error)
}

type BlockDatabase interface {
	AddRule(*block.Rule) error
	RemoveRule(*block.Rule) error
	IsBlocked(string, int) (bool, error)
	ListRules(context.Context, ListBlockRulesFunc)
}
