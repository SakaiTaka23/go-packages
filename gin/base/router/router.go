package router

import (
	"base/handler"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func SetRouter(r *gin.Engine) *gin.Engine {
	useLogger(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	handler.NewUserHandler(r)

	return r
}

func useLogger(r *gin.Engine) *gin.Engine {
	logger, _ := zap.NewProduction()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	return r
}
