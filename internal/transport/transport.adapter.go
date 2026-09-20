package transport

import "net/http"

func Adapt(f HandlerFn) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			w.WriteHeader(100)
			w.Write([]byte(err.Error()))
		}
	})
}

func MuxAdapt(mux *http.ServeMux, pattern string, f HandlerFn) {
	mux.Handle(pattern, Adapt(f))
}
