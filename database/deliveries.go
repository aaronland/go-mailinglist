package database

import (
	"context"

	"github.com/aaronland/go-mailinglist/v2/delivery"
)

type ListDeliveriesFunc func(*delivery.Delivery) error

type DeliveriesDatabase interface {
	AddDelivery(context.Context, *delivery.Delivery) error
	ListDeliveries(context.Context, ListDeliveriesFunc) error
	GetDeliveryWithAddressAndMessageId(context.Context, string, string) (*delivery.Delivery, error)
	Close() error
}
