package router

import (
	"base/handler"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"time"
)

type middleware struct {
	l *zap.Logger
}

func SetRouter(r *gin.Engine) *gin.Engine {
	logger, _ := zap.NewProduction()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	midd := initMiddleware()
	r.Use(midd.useLogger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	handler.NewUserHandler(r)

	return r
}

func initMiddleware() *middleware {
	return &middleware{}
}

func (m *middleware) useLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Print("request start: ", c.Request.URL)
		c.Next()
	}
}
