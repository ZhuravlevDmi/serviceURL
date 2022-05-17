package mymiddleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
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

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	// w.Writer будет отвечать за gzip-сжатие, поэтому пишем в него
	return w.Writer.Write(b)
}

func GzipHandle(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Acept-Encoding"), "gzip") {
			// если gzip не поддерживается, передаём управление
			// дальше без изменений
			h.ServeHTTP(w, r)
			return
		}
		// создаем gzip.Writer поверх текущего w
		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		h.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
	}
	return http.HandlerFunc(fn)
}
