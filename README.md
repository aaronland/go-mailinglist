# go-mailinglist

Minimalist and opinionated scaffolding for running a mailing list.

## Documentation

Documentation is incomplete at this time.

## Motivation (and caveats)

I needed "the simplest and dumbest thing" to manage delivering email messages to a list of subscribers which I could run locally _and_ using "scale-to-zero" cloud services. This is the second version of that thing.

It does very little and all of it from the command line (or equivalent). For example it does not provide a web interface for people to subscribe or unsubscribe to the mailing list. Most of the code to run that web application has been written but it still needs more polish and has not been the priority yet.

## Databases

This package defines a number of Go language interfaces for managing mailing list related database tables (subscribers, deliveries, event logs). Currently this one implementation of those interfaces that uses the [gocloud.dev/docstore](https://pkg.go.dev/gocloud.dev/docstore) abstraction layer for storing and retrieving data from a document store.

So far this package has been tested with and enables support for [DynamoDB](https://gocloud.dev/howto/docstore/#dynamodb) by default.

## Sending email

This package uses the [aaronland/gomail-sender](https://github.com/aaronland/gomail-sender) package for delivering email messages using a variety of "providers".

So far this package has been tested with and enables support for [Amazon Simple Email Service (SES)[https://github.com/aaronland/gomail-sender-ses] by default.

## Attachments

Message attachments are loaded using the [gocloud.dev/blob](https://pkg.go.dev/gocloud.dev/blob) abstraction layer for reading files from a variety of local or remote document storage systems.

So far this package has been tested with and enables support for loading attachments from the [local filesystem](https://gocloud.dev/howto/blob/#local) and the [Amazon Simple Storage Service (S3)](https://gocloud.dev/howto/blob/#s3).

## Example

The following example assumes a [local instance of DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.DownloadingAndRunning.html)  running from a [Docker](https://www.docker.com/) container and the command line tools described below. For examepl:

```
$> docker run --rm -it -p 8000:8000 amazon/dynamodb-local
WARNING: The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no specific platform was requested
Initializing DynamoDB Local with the following configuration:
Port:	8000
InMemory:	true
DbPath:	null
SharedDb:	false
shouldDelayTransientStatuses:	false
CorsParams:	null
```

The first step is to create the database tables necessary for the mailing list. Use the `local-tables` Makefile target to do this:

```
$> make local-tables
go run -mod vendor -ldflags="-s -w" cmd/create-dynamodb-tables/main.go \
		-refresh \
		-client-uri 'aws://?region=localhost&credentials=anon:&local=true'
```

Next, add a subscriber to the mailing list. Use the `local-add` Makefile target to do this:

```
$> make local-add ADDRESS=bob@localhost
go run -mod vendor -ldflags="-s -w" cmd/add-subscriptions/main.go \
		-subscriptions-database-uri 'awsdynamodb://subscriptions?region=localhost&credentials=anon:&local=true&partition_key=Address&allow_scans=true' \
		-address bob@localhost \
		-confirmed
{"address":"bob@localhost","created":1732470161,"confirmed":1732470161,"lastmodified":1732470161,"status":1}
2024/11/24 09:42:41 INFO Subscription added address=bob@localhost
```

You can confirm the subscriber has been added by using the `local-list` Makefile target:

```
$> make local-list
go run -mod vendor -ldflags="-s -w" cmd/list-subscriptions/main.go \
		-subscriptions-database-uri 'awsdynamodb://subscriptions?region=localhost&credentials=anon:&local=true&partition_key=Address&allow_scans=true'
{"address":"bob@localhost","created":1732470161,"confirmed":1732470161,"lastmodified":1732470161,"status":1}
```

Finally, deliver a message (with an attachment) to all the subscribers using the `local-deliver` Makefile target:

```
$> make local-deliver
go run -mod vendor -ldflags="-s -w" cmd/deliver-message/main.go \
		-subscriptions-database-uri 'awsdynamodb://subscriptions?region=localhost&credentials=anon:&local=true&partition_key=Address&allow_scans=true' \
		-deliveries-database-uri 'awsdynamodb://deliveries?region=localhost&credentials=anon:&local=true&partition_key=Address&allow_scans=true' \
		-eventlogs-database-uri 'awsdynamodb://eventlogs?region=localhost&credentials=anon:&local=true&partition_key=Address&allow_scans=true' \
		-sender-uri stdout:// \
		-from do-not-reply@localhost \
		-subject  \
		-body  \
		-attachment file:///Users/asc/code/go-mailinglist/fixtures/hellokitty.jpg
Mime-Version: 1.0
Date: Sun, 24 Nov 2024 09:44:11 -0800
X-MailingList-Id: 057c4bf7-f8b1-41a4-be10-fbfd9ca85088
From: <do-not-reply@localhost>
To: <bob@localhost>
Subject: -body
Content-Type: multipart/related;
 boundary=5742b7f14e3b26c910ce1b93411111df97a01bd92f5a3708d956d1f9c6e7

--5742b7f14e3b26c910ce1b93411111df97a01bd92f5a3708d956d1f9c6e7
Content-Transfer-Encoding: quoted-printable
Content-Type: text/plain; charset=UTF-8


--5742b7f14e3b26c910ce1b93411111df97a01bd92f5a3708d956d1f9c6e7
Content-Disposition: inline; filename="hellokitty.jpg"
Content-ID: <hellokitty.jpg>
Content-Transfer-Encoding: base64
Content-Type: image/jpeg; name="hellokitty.jpg"

/9j/4AAQSkZJRgABAQEAkACQAAD/2wBDAAcFBQYFBAcGBgYIBwcICxILCwoKCxYPEA0SGhYbGhkW
GRgcICgiHB4mHhgZIzAkJiorLS4tGyIyNTEsNSgsLSz/2wBDAQcICAsJCxULCxUsHRkdLCwsLCws
... and so on
```

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
