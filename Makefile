CWD=$(shell pwd)

devel:
	go run -mod vendor cmd/subscriptiond/main.go -devel -templates 'templates/html/*.html'

tools:
	go build -mod vendor -o bin/subscribe cmd/subscribe/main.go
	go build -mod vendor -o bin/subscriptiond cmd/subscriptiond/main.go
