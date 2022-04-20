package main

import (
	"github.com/ZhuravlevDmi/serviceURL/internal/config"
	"github.com/ZhuravlevDmi/serviceURL/internal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.MainHandler)
	http.ListenAndServe(config.Port, nil)
}
