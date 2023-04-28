package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PublishError(c *gin.Context, err error) {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
