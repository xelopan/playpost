package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (UserController) GetRecommended(c *gin.Context) {
	// GET returns list of recommended users to be added as a friend. A complicated algorithm
	// and parameters might be needed, but we don't care about that for now. Just show all users
	c.String(http.StatusOK, "Get list of recommended user")
}

func (UserController) GetDetails(c *gin.Context) {
	// GET get details of a user, and also status of friend request
	c.String(http.StatusOK, "Details of a user")
}

func (UserController) Search(c *gin.Context) {
	// GET get the list of user
	c.String(http.StatusOK, "Searching user by name or handle")
}

func (UserController) GetFriends(c *gin.Context) {
	// GET friend list of a user (also friend requests to or by the user)
	c.String(http.StatusOK, "Friend list of a user")
}

func (UserController) CreateFriendRequest(c *gin.Context) {
	// POST
	c.String(http.StatusOK, "Creating a friend request")
}

func (UserController) RespondFriendRequest(c *gin.Context) {
	// PATCH
	c.String(http.StatusOK, "Respond to a friend request")
}

func (UserController) CancelFriendRequest(c *gin.Context) {
	// DELETE
	c.String(http.StatusOK, "Cancel a friend request")
}
