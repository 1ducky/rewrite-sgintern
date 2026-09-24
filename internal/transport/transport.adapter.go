package transport

import (
	"RewriteProject/internal/app/err"
	"RewriteProject/internal/transport/middleware/tracer"
	"encoding/json"
	"log"
	"net/http"
)

type Adapter struct {
	errApp *err.ErrApp
}

func NewAdapter(errApp *err.ErrApp) *Adapter {
	return &Adapter{
		errApp: errApp,
	}
}

func (a *Adapter) Adapt(f HandlerFn, domain err.ErrDomain) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if errReq := f(w, r); errReq != nil {
			entry := a.errApp.Translate(errReq, domain)
			if entry.Code == err.InternalErrCode {
				log.Print("Invariant Detected : ", errReq, " requestid: ", tracer.GetTraceID(r.Context()))
			}
			w.WriteHeader(entry.StatusCode)
			json.NewEncoder(w).Encode(ErrorResponse{
				Code:    string(entry.Code),
				Message: entry.Message,
				TraceID: tracer.GetTraceID(r.Context()),
				Status:  StatusError,
			})
		}

	}
}
