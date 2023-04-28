package controllers

import (
	"net/http"
	"playpost/core/app/models"
	"playpost/core/app/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

// Attempt to authenticate using supplied credentials.
func (AuthController) Auth(c *gin.Context) {
	// TODO implement a way to authenticate user
	c.String(http.StatusOK, "Authenticate!")
}

func (AuthController) Register(c *gin.Context) {
	authSvc := new(services.AuthService)
	var err error

	var reg models.Register
	c.Request.ParseForm()
	err = c.Bind(&reg)
	if err != nil {
		PublishError(c, err)
		return
	}
	// userdata := make(map[string]string)

	_, err = authSvc.Register(&reg)
	if err != nil {
		PublishError(c, err)
		return
	}
	c.String(http.StatusOK, "User registered, try to login!")
}

func (AuthController) Logout(c *gin.Context) {
	// TODO implement a way to logout a user
	c.String(http.StatusOK, "Logout!")
}

func (AuthController) ResetPassword(c *gin.Context) {
	// TODO implement a way to reset user's password
	c.String(http.StatusOK, "Password Reset!")
}
