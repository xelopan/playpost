package main

import (
	"net/http"

	"playpost/core/api"
)

func main() {
	http.HandleFunc("/", api.Handler)
	http.ListenAndServe("localhost:8080", nil)
}
