package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

func (PostController) CreatePost(c *gin.Context) {
	c.String(http.StatusOK, "Creating a post!")
}

func (PostController) GetPost(c *gin.Context) {
	c.String(http.StatusOK, "Getting Post Details!")
}
