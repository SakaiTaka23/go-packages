package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct{}

func NewUserHandler(r *gin.Engine) {
	handler := &UserHandler{}
	user := r.Group("/user")
	{
		user.GET("", handler.getUserMessage())
		user.GET("/:name", handler.getUser)
	}
}

// use this if it is a middleware
func (u *UserHandler) getUserMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "user",
		})
	}
}

// use this style for handlers in general
func (u *UserHandler) getUser(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hi " + name,
	})
}
