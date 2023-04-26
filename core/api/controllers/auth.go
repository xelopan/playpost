package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

// Attempt to authenticate using supplied credentials.
func (AuthController) Auth(c *gin.Context) {
	// TODO implement a way to authenticate user
	c.String(http.StatusOK, "Authenticate!")
}

func (AuthController) Register(c *gin.Context) {
	// TODO implement a way to register a user
	c.String(http.StatusOK, "Register!")
}

func (AuthController) Logout(c *gin.Context) {
	// TODO implement a way to logout a user
	c.String(http.StatusOK, "Logout!")
}

func (AuthController) ResetPassword(c *gin.Context) {
	// TODO implement a way to reset user's password
	c.String(http.StatusOK, "Password Reset!")
}
