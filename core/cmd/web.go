package main

import (
	"net/http"

	"github.com/xelopan/playpost-core/api"
)

func main() {
	http.HandleFunc("/", api.Handler)
	http.ListenAndServe("localhost:8080", nil)
}
