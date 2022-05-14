package main

import (
	"github.com/ZhuravlevDmi/serviceURL/internal/config"
	"github.com/ZhuravlevDmi/serviceURL/internal/handlers"
	"github.com/ZhuravlevDmi/serviceURL/internal/mymiddleware"
	"github.com/ZhuravlevDmi/serviceURL/internal/storage"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

func main() {
	var cfgAdr config.ConfigAdress
	cfgAdr.Parse()

	var MapURL storage.Storage = &storage.StorageMapURL{MapURL: make(map[string]string)}

	r := chi.NewRouter()

	// зададим встроенные Middleware, чтобы улучшить стабильность приложения
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// middleware котороый пропускает только post или get запрос
	r.Use(mymiddleware.MethodRequestMiddleware)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/", func(r chi.Router) {
		r.Post("/api/shorten", handlers.HandlerAPIShorten(MapURL, cfgAdr.BaseURL))
		r.Get("/{path}", handlers.HandlerGetURL(MapURL, cfgAdr.BaseURL))
		r.Post("/", handlers.HandlerPostURL(MapURL, cfgAdr.BaseURL))
	})

	http.ListenAndServe(cfgAdr.ServerAddress, r)
}
