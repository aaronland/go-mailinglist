CWD=$(shell pwd)

GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

LOCAL_SUBSCRIPTIONS_DATABASE_URI=awsdynamodb://subscriptions?region=localhost&credentials=anon:&local=true&partition_key=Address&allow_scans=true
LOCAL_DELIVERIES_DATABASE_URI=awsdynamodb://deliveries?region=localhost&credentials=anon:&local=true&partition_key=Address&allow_scans=true
LOCAL_EVENTLOGS_DATABASE_URI=awsdynamodb://eventlogs?region=localhost&credentials=anon:&local=true&partition_key=Address&allow_scans=true

LOCAL_SENDER_URI=stdout://
LOCAL_DELIVER_FROM=do-not-reply@localhost
LOCAL_ATTACHMENT=file://$(CWD)/fixtures/hellokitty.jpg

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
		-subscriptions-database-uri '$(LOCAL_SUBSCRIPTIONS_DATABASE_URI)' \
		-address $(ADDRESS) \
		-confirmed

local-status:
	go run -mod $(GOMOD) -ldflags="-s -w" cmd/set-subscription-status/main.go \
		-subscriptions-database-uri '$(LOCAL_SUBSCRIPTIONS_DATABASE_URI)' \
		-address $(ADDRESS) \
		-status $(STATUS)

local-list:
	go run -mod $(GOMOD) -ldflags="-s -w" cmd/list-subscriptions/main.go \
		-subscriptions-database-uri '$(LOCAL_SUBSCRIPTIONS_DATABASE_URI)'

local-deliver:
	go run -mod $(GOMOD) -ldflags="-s -w" cmd/deliver-message/main.go \
		-subscriptions-database-uri '$(LOCAL_SUBSCRIPTIONS_DATABASE_URI)' \
		-deliveries-database-uri '$(LOCAL_DELIVERIES_DATABASE_URI)' \
		-eventlogs-database-uri '$(LOCAL_EVENTLOGS_DATABASE_URI)' \
		-sender-uri $(LOCAL_SENDER_URI) \
		-from $(LOCAL_DELIVER_FROM) \
		-subject $(SUBJECT) \
		-body $(BODY) \
		-attachment $(LOCAL_ATTACHMENT)
