# go-mailinglist

Minimalist and opinionated scaffolding for running a mailing list.

## Documentation

Documentation is incomplete at this time.

## Tools

```
$> make cli
go build -mod vendor -ldflags="-s -w"  -o bin/create-dynamodb-tables cmd/create-dynamodb-tables/main.go
go build -mod vendor -ldflags="-s -w"  -o bin/add-subscriptions cmd/add-subscriptions/main.go
go build -mod vendor -ldflags="-s -w"  -o bin/remove-subscriptions cmd/remove-subscriptions/main.go
go build -mod vendor -ldflags="-s -w"  -o bin/status-subscriptions cmd/set-subscription-status/main.go
go build -mod vendor -ldflags="-s -w"  -o bin/list-subscriptions cmd/list-subscriptions/main.go
```

### create-dynamodb-tables

### add-subscriptions

### remove-subscriptions

### set-subscruption-status

### list-subscriptions

### send-message

## See also

* https://github.com/aaronland/gomail
* https://github.com/aaronland/gomail-ses
