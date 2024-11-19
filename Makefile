CWD=$(shell pwd)

GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)"  -o bin/list-subscriptions cmd/list-subscriptions/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)"  -o bin/add-subscriptions cmd/add-subscriptions/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)"  -o bin/remove-subscriptions cmd/remove-subscriptions/main.go
