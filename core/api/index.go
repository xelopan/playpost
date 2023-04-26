package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

// As the Handler func is called by vercel, we initialize the router for this API.
func init() {
	app = NewRouter()
}

// This function will be the entrypoint used by vercel.
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
