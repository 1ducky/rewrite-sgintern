package transport

import (
	"context"
	"log/slog"
	"net/http"
)

type TransportAPI interface {
	Start() error
	Stop(context.Context) error
}

type Transport struct {
	server  *http.Server
	log     *slog.Logger
	Handler http.Handler
}
