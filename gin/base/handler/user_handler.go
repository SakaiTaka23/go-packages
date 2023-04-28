package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UserHandler struct{}

func NewUserHandler(r *gin.Engine) {
	handler := &UserHandler{}
	user := r.Group("/user")
	{
		user.GET("", handler.index())
		user.GET("/:name", handler.show)
		user.POST("", handler.create)
	}
}

// use this if it is a middleware
func (u *UserHandler) index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "user",
		})
	}
}

// use this style for handlers in general
func (u *UserHandler) show(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hi " + name,
	})
}

func (u *UserHandler) create(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "validation failed",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "validation passed",
		"user":    &user,
	})
}
