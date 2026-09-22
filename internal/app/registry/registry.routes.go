package registry

import (
	"RewriteProject/internal/app/err"
	"RewriteProject/internal/app/services/auth"
	"RewriteProject/internal/transport"
	"net/http"
)

func NewRegisterRoutes(handlers RegisterHandler, adapter *transport.Adapter) http.Handler {
	mux := http.NewServeMux()
	routeAuth(mux, handlers.AuthHandler, adapter)
	return mux
}

func routeAuth(mux *http.ServeMux, handler auth.Handler, adapter *transport.Adapter) {
	mux.HandleFunc("POST /auth/register", adapter.Adapt(handler.Register, err.Auth))
	mux.HandleFunc("POST /auth/login", adapter.Adapt(handler.Login, err.Auth))
}
