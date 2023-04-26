package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LikeController struct{}

func (LikeController) LikePost(c *gin.Context) {
	c.String(http.StatusOK, "Liking a post")
}

func (LikeController) LikeComment(c *gin.Context) {
	c.String(http.StatusOK, "Liking a comment")
}
