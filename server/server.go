package server

import (
	"errors"
	"fmt"
	_ "log"
	"net/http"
	"net/url"
	"strings"
)

type Server interface {
	ListenAndServe(*http.ServeMux) error
	Address() string
}

func NewServer(protocol string, host string, port int, args ...interface{}) (Server, error) {

	address := fmt.Sprintf("http://%s:%d", host, port)

	u, err := url.Parse(address)

	if err != nil {
		return nil, err
	}

	var svr Server

	switch strings.ToUpper(protocol) {

	case "HTTP":

		svr, err = NewHTTPServer(u, args...)

	case "LAMBDA":

		svr, err = NewLambdaServer(u, args...)

	default:
		return nil, errors.New("Invalid server protocol")
	}

	if err != nil {
		return nil, err
	}

	return svr, nil
}
