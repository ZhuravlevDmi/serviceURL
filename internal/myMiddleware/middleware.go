package myMiddleware

import (
	"net/http"
)

func MethodRequestMiddleware(h http.Handler) http.Handler {
	// middleware пропускает только GET или POST запрос
	fn := func(w http.ResponseWriter, r *http.Request) {
		if method := r.Method; method != http.MethodGet && method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Обрабатываются только GET или POST запрос"))
			return
		}
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
