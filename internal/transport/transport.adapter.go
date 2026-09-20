package transport

import (
	"RewriteProject/internal/app/err"
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

func (a *Adapter) Adapt(f HandlerFn, domain err.ErrDomain) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			entry := a.errApp.Translate(err, domain)
			w.WriteHeader(entry.StatusCode)
			w.Write([]byte(entry.Message))
		}
	})
}
