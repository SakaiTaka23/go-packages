package router

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func SetRouter(r *gin.Engine) *gin.Engine {
	useLogger(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userRoute(r)

	return r
}

func useLogger(r *gin.Engine) *gin.Engine {
	logger, _ := zap.NewProduction()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	return r
}

func userRoute(r *gin.Engine) *gin.Engine {
	user := r.Group("/user")
	{
		user.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "user",
			})
		})
	}

	return r
}
