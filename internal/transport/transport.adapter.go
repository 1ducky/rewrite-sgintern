package transport

import (
	"RewriteProject/internal/app/err"
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
		if err := f(w, r); err != nil {
			log.Print(err)
			entry := a.errApp.Translate(err, domain)
			w.WriteHeader(entry.StatusCode)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    entry.Code,
				"message": entry.Message,
			})
		}

	}
}
