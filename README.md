# go-mailinglist

Minimalist and opinionated scaffolding for running a mailing list.

## Documentation

Documentation is incomplete at this time.

## Motivation (and caveats)

I needed "the simplest and dumbest thing" to manage delivering email messages to a list of subscribers which I could run locally _and_ using "scale-to-zero" cloud services. This is the second version of that thing.

It does very little and all of it from the command line (or equivalent). For example it does not provide a web interface for people to subscribe or unsubscribe to the mailing list. Most of the code to run that web application has been written but it still needs more polish and has not been the priority yet.

## Databases

This package defines a number of Go language interfaces for managing mailing list related database tables (subscribers, deliveries, event logs). Currently this one implementation of those interfaces that uses the [gocloud.dev/docstore](https://pkg.go.dev/gocloud.dev/docstore) abstraction layer for storing and retrieving data from a document store.

These databases are currently designed only support a single mailing list.

So far this package has been tested with and enables support for [DynamoDB](https://gocloud.dev/howto/docstore/#dynamodb) by default.

## Sending email

This package uses the [aaronland/gomail-sender](https://github.com/aaronland/gomail-sender) package for delivering email messages using a variety of "providers".

So far this package has been tested with and enables support for [Amazon Simple Email Service (SES)[https://github.com/aaronland/gomail-sender-ses] by default.

## Attachments

Message attachments are loaded using the [gocloud.dev/blob](https://pkg.go.dev/gocloud.dev/blob) abstraction layer for reading files from a variety of local or remote document storage systems.

So far this package has been tested with and enables support for loading attachments from the [local filesystem](https://gocloud.dev/howto/blob/#local) and the [Amazon Simple Storage Service (S3)](https://gocloud.dev/howto/blob/#s3).

## Example

The following example assumes a [local instance of DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.DownloadingAndRunning.html)  running from a [Docker](https://www.docker.com/) container and the command line tools described below. For example:

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
		-attachment /usr/local/go-mailinglist/fixtures/hellokitty.jpg
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
go build -mod vendor -ldflags="-s -w"  -o bin/list-subscriptions cmd/list-subscriptions/main.go
go build -mod vendor -ldflags="-s -w"  -o bin/add-subscriptions cmd/add-subscriptions/main.go
go build -mod vendor -ldflags="-s -w"  -o bin/remove-subscriptions cmd/remove-subscriptions/main.go
go build -mod vendor -ldflags="-s -w"  -o bin/set-subscription-status cmd/set-subscription-status/main.go
go build -mod vendor -ldflags="-s -w"  -o bin/create-dynamodb-tables cmd/create-dynamodb-tables/main.go
go build -mod vendor -ldflags="-s -w"  -o bin/deliver-mail cmd/deliver-message/main.go
```

### create-dynamodb-tables

Instantiate a new set of DynamoDB tables for a mailing list.

```
$> ./bin/create-dynamodb-tables -h
Usage of ./bin/create-dynamodb-tables:
  -client-uri string
    	A registered aaronland/go-aws-dynamodb URI. (default "aws://?region=localhost&credentials=anon:&local=true")
  -prefix string
    	Optional string to prepend to all table names.
  -refresh
    	Refresh (delete and recreate) tables that have already been created.
```

### add-subscriptions

Add a subscriber to the mailing list

```
$> ./bin/add-subscriptions -h
  -address value
    	One or more addresses to add as subscriptions
  -confirmed
    	A boolean flag indicating whether the subscriber is confirmed.
  -subscriptions-database-uri string
    	A registered aaronland/go-mailinglist/v2/database.SubscriptionsDatabase URI.
  -verbose
    	Enable verbose (debug) logging.
```

### remove-subscriptions

Remove a subscriber from the mailing list.

```
$> ./bin/remove-subscriptions -h
  -address value
    	One or more addresses whose subscriptions should be removed.
  -subscriptions-database-uri string
    	A registered aaronland/go-mailinglist/v2/database.SubscriptionsDatabase URI.
  -verbose
    	Enable verbose (debug) logging.
```

### set-subscription-status

Assign the subscription status for one or more subscribers.

```
$> ./bin/set-subscription-status -h
  -address value
    	One or more addresses whose subscriptions should be removed.
  -status string
    	The status to assign to each address.
  -subscriptions-database-uri string
    	A registered aaronland/go-mailinglist/v2/database.SubscriptionsDatabase URI.
  -verbose
    	Enable verbose (debug) logging.
```

### list-subscriptions

List all of the subscribers for a mailing list.

```
$> ./bin/list-subscriptions -h
  -subscriptions-database-uri string
    	A registered aaronland/go-mailinglist/v2/database.SubscriptionsDatabase URI.
  -verbose
    	Enable verbose (debug) logging.
```

### deliver-message

Deliver a message to one or more addresses.

```
$> ./bin/deliver-mail -h
  -attachment value
    	Zero or more URIs referencing files to attach to the message. URIs are dereferenced using the aaronland/gocloud-blob/bucket.ParseURI method.
  -body string
    	The body of the message being delivered. If "-" then body will be read from STDIN.
  -content-type string
    	The content-type of the message body. (default "text/plain")
  -deliveries-database-uri string
    	A registered aaronland/go-mailinglist/v2/database.DeliveriesDatabase URI.
  -eventlogs-database-uri string
    	A registered aaronland/go-mailinglist/v2/database.EventLogsDatabase URI.
  -from string
    	The address delivering the message.
  -message-id string
    	Optional custom message ID to assign to the message. If empty a unique key will be generated on delivery.
  -sender-uri string
    	A registered aaronland/go-mail.Sender URI.
  -subject string
    	The subject of the message being delivered.
  -subscriptions-database-uri string
    	A registered aaronland/go-mailinglist/v2/database.SubscriptionsDatabase URI.
  -to value
    	One or more addresses to deliver the message to. If empty then the message will be delivered to all subscribers whose status is "enabled".
  -verbose
    	Enable verbose (debug) logging.
```

## AWS

### DynamoDB

The following URI schemes are available for instantiating DynamoDB database connections:

* `dynamodb://` – the default scheme supported by the [gocloud.dev/docstore](https://pkg.go.dev/gocloud.dev/docstore) package.
* `awsdynamodb://` – and alternate scheme that supports defining AWS credentials using labeled strings, described below.

### S3

The following URI schemes are available for instantiating S3 data storage instances:

* `s3://` – the default scheme supported by the [gocloud.dev/blob](https://pkg.go.dev/gocloud.dev/blob) package.
* `s3blob://` – and alternate scheme that supports defining AWS credentials using labeled strings, described below.

### Credentials

Credentials strings (as `?credentials={CREDENTIALS}`) are expected to be a valid [aaronland/go-aws-auth](https://github.com/aaronland/go-aws-auth?tab=readme-ov-file#credentials) credentials "label":

| Label | Description |
| --- | --- |
| `anon:` | Empty or anonymous credentials. |
| `env:` | Read credentials from AWS defined environment variables. |
| `iam:` | Assume AWS IAM credentials are in effect. |
| `sts:{ARN}` | Assume the role defined by `{ARN}` using STS credentials. |
| `{AWS_PROFILE_NAME}` | This this profile from the default AWS credentials location. |
| `{AWS_CREDENTIALS_PATH}:{AWS_PROFILE_NAME}` | This this profile from a user-defined AWS credentials location. |

For example:

```
aws:///us-east-1?credentials=iam:
```

## See also

* https://github.com/aaronland/gomail
* https://github.com/aaronland/gomail-ses
