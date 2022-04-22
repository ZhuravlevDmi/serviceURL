package main

import (
	"github.com/ZhuravlevDmi/serviceURL/internal/handlers"
	"github.com/ZhuravlevDmi/serviceURL/internal/storage"
	"net/http"
)

func main() {
	var MapURL storage.Storage = &storage.StorageMapURL{MapURL: map[string]string{}}
	http.HandleFunc("/", handlers.MainHandler(MapURL))
	http.ListenAndServe(":8080", nil)
}
