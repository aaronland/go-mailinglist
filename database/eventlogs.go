package database

import (
	"context"

	"github.com/aaronland/go-mailinglist/v2/eventlog"	
)

type ListEventLogsFunc func(*eventlog.EventLog) error

type EventLogsDatabase interface {
	AddEventLog(context.Context, *eventlog.EventLog) error
	ListEventLogs(context.Context, ListEventLogsFunc) error
	Close() error
}
