package main

import (
	"github.com/ZhuravlevDmi/serviceURL/internal/handlers"
	"github.com/ZhuravlevDmi/serviceURL/internal/myMiddleware"
	"github.com/ZhuravlevDmi/serviceURL/internal/storage"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

func main() {
	var MapURL storage.Storage = &storage.StorageMapURL{MapURL: make(map[string]string)}

	r := chi.NewRouter()

	// зададим встроенные Middleware, чтобы улучшить стабильность приложения
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// middleware котороый пропускает только post или get запрос
	r.Use(myMiddleware.MethodRequestMiddleware)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/", func(r chi.Router) {
		r.Get("/{path}", handlers.HandlerGetURL(MapURL))
		r.Post("/", handlers.HandlerPostURL(MapURL))

	})

	http.ListenAndServe(":8080", r)
}
