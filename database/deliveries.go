package database

import (
	"context"
	"github.com/aaronland/go-mailinglist/delivery"
)

type DeliveriesDatabase interface {
	AddDelivery(*delivery.Delivery) error
	ListDeliveries(context.Context, ListDeliveriesFunc) error
	GetDeliveryWithAddressAndMessageId(string, string) (*delivery.Delivery, error)
}
