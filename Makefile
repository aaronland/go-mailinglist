CWD=$(shell pwd)

GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)"  -o bin/list-subscriptions cmd/list-subscriptions/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)"  -o bin/add-subscriptions cmd/add-subscriptions/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)"  -o bin/remove-subscriptions cmd/remove-subscriptions/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)"  -o bin/status-subscriptions cmd/set-subscription-status/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)"  -o bin/create-dynamodb-tables cmd/create-dynamodb-tables/main.go

# docker run --rm -it -p 8000:8000 amazon/dynamodb-local

local-tables:
	go run -mod $(GOMOD) -ldflags="-s -w" cmd/create-dynamodb-tables/main.go \
		-refresh \
		-client-uri 'aws://?region=localhost&credentials=anon:&local=true'

local-add:
	go run -mod $(GOMOD) -ldflags="-s -w" cmd/add-subscriptions/main.go \
		-subscriptions-database-uri 'awsdynamodb://subscriptions?region=localhost&credentials=anon:&local=true&partition_key=Address&allow_scans=true' \
		-address $(ADDRESS) \
		-confirmed

local-status:
	go run -mod $(GOMOD) -ldflags="-s -w" cmd/set-subscription-status/main.go \
		-subscriptions-database-uri 'awsdynamodb://subscriptions?region=localhost&credentials=anon:&local=true&partition_key=Address&allow_scans=true' \
		-address $(ADDRESS) \
		-status $(STATUS)

local-list:
	go run -mod $(GOMOD) -ldflags="-s -w" cmd/list-subscriptions/main.go \
		-subscriptions-database-uri 'awsdynamodb://subscriptions?region=localhost&credentials=anon:&local=true'

local-deliver:
	go run -mod $(GOMOD) -ldflags="-s -w" cmd/deliver-message/main.go \
		-subscriptions-database-uri 'awsdynamodb://subscriptions?region=localhost&credentials=anon:&local=true&partition_key=Address&allow_scans=true' \
		-deliveries-database-uri 'awsdynamodb://deliveries?region=localhost&credentials=anon:&local=true&partition_key=Address&allow_scans=true' \
		-eventlogs-database-uri 'awsdynamodb://eventlogs?region=localhost&credentials=anon:&local=true&partition_key=Address&allow_scans=true' \
		-sender-uri stdout:// \
		-from do-not-reply@localhost \
		-to $(TO) \
		-subject $(SUBJECT) \
		-body $(BODY) \
		-attachment $(CWD)/fixtures/hellokitty.jpg
