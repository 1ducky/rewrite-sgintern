package transport

import (
	"context"
	"net/http"
)

type TransportAPI interface {
	Start() error
	Stop(context.Context) error
	Address() string
}

type Transport struct {
	server  *http.Server
	Handler http.Handler
}

type HandlerFn func(w http.ResponseWriter, r *http.Request) error
