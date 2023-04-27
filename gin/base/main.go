package main

import (
	"base/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	router.SetRouter(r)
	_ = r.Run("127.0.0.1:8080")
}
