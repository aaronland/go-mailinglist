CWD=$(shell pwd)

GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)"  -o bin/list-subscriptions cmd/list-subscriptions/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)"  -o bin/add-subscriptions cmd/add-subscriptions/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)"  -o bin/remove-subscriptions cmd/remove-subscriptions/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)"  -o bin/create-dynamodb-tables cmd/create-dynamodb-tables/main.go

# docker run --rm -it -p 8000:8000 amazon/dynamodb-local

local-tables:
	go run -mod $(GOMOD) -ldflags="-s -w" cmd/create-dynamodb-tables/main.go \
		-refresh \
		-client-uri 'aws://?region=localhost&credentials=anon:&local=true'

local-add:
	go run -mod $(GOMOD) -ldflags="-s -w" cmd/add-subscriptions/main.go \
		-subscriptions-database-uri 'awsdynamodb://subscriptions?region=localhost&credentials=anon:&local=true' \
		-address $(ADDRESS) \
		-confirmed
