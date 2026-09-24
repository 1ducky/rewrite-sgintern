package middleware

import (
	"RewriteProject/internal/transport"
	"RewriteProject/internal/transport/middleware/tracer"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func TrackerMiddleware() func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			// TODO : genorate request id and add to context
			requestId := uuid.New().String()
			ctxtracer := tracer.WithTracer(r.Context(), requestId)
			r = r.WithContext(ctxtracer)
			tw := &transport.TrasnportWriter{ResponseWriter: w}

			// TODO : add request id to response header
			w.Header().Set(tracer.TraceIDKey, requestId)
			// TODO : add request id to logger
			next.ServeHTTP(tw, r)
			endReq := time.Since(startTime)

			log.Print("requestid: ", requestId, "", " [", tw.StatusCode, "]  path: ", r.URL.Path, " method: ", r.Method, " cost: ", endReq)
		})
	}
}
