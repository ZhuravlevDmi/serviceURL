package main

import (
	"github.com/ZhuravlevDmi/serviceURL/internal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.MainHandler)
	http.ListenAndServe(":8080", nil)
}
