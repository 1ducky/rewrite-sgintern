package middleware

import (
	"RewriteProject/internal/transport"
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func TrackerMiddleware(log *slog.Logger) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			// TODO : genorate request id and add to context
			requestId := uuid.New().String()
			ctx := context.WithValue(r.Context(), "requestid", requestId)
			r = r.WithContext(ctx)
			tw := &transport.TrasnportWriter{ResponseWriter: w}

			// TODO : add request id to response header
			w.Header().Set("requestid", requestId)
			// TODO : add request id to logger
			next.ServeHTTP(tw, r)
			endReq := time.Since(startTime)

			log.Info("requestid: ", requestId, "", "[", tw.StatusCode, "]  url: ", r.URL.Path, "method: ", r.Method, "", endReq)
		})
	}
}
