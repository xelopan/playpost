package middlewares

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	//TODO implement the logic to prevent access if user is not authenticated
	// (or just use a pre-made package!)
	return func(c *gin.Context) {
		c.Next()
	}
}
