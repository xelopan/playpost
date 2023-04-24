package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func Handler(w http.ResponseWriter, r *http.Request) {
	// entry point for all http request
	app.ServeHTTP(w, r)
}

func route(r *gin.RouterGroup) {
	r.GET("/admin", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from golang in vercel!")
	})
}

func init() {
	app = gin.New()
	r := app.Group("/core/api")
	route(r)
}
