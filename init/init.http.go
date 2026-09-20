package init

import (
	"context"
	"log/slog"
	"net/http"

	"RewriteProject/internal/config"
	"RewriteProject/internal/transport"
	"RewriteProject/internal/transport/middleware"
)

func InitialiseHTTP(ctx context.Context, config config.HTTPConfig, routes http.Handler, log *slog.Logger) transport.TransportAPI {
	handler := middleware.Chain(routes, middleware.TrackerMiddleware(log))
	return transport.NewTransport(config, handler, log)
}

// func SetupRoute(ctx context.Context, ucs *ucs, cfg config.AppConfig, router *http.ServeMux) http.Handler {

// 	w := chi.NewMux()
// 	// setup middleware
// 	w.Use(cors.Handler(cors.Options{
// 		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8000"},
// 		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
// 		AllowedHeaders:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
// 		ExposedHeaders:   []string{"Content-Length"},
// 		AllowCredentials: true,
// 		MaxAge:           300,
// 	}))
// 	w.Use(middleware.RecoverMiddleware(ucs.assets.Logger))
// 	w.Use(middleware.LimiterMiddleware(cfg.RateLimitConfig, ucs.assets.Logger))

// 	// Setup router
// 	route := route{
// 		mux:  router,
// 		logs: ucs.assets.Logger,
// 	}
// 	route.setup()    // load route
// 	route.fallback() // 404 and 500 handler

// 	return w
// }
