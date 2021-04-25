package database

import (
	"context"
	"github.com/aaronland/go-mailinglist/eventlog"
)

type EventLogsDatabase interface {
	AddEventLog(*eventlog.EventLog) error
	ListEventLogs(context.Context, ListEventLogsFunc) error
}
