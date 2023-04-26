package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct{}

func (CommentController) CommentOnPost(c *gin.Context) {
	c.String(http.StatusOK, "Send a comment on a post")
}

func (CommentController) ReplyComment(c *gin.Context) {
	c.String(http.StatusOK, "Reply on another comment")
}

func (CommentController) DeleteComment(c *gin.Context) {
	c.String(http.StatusOK, "Delete a comment")
}
