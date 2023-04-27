package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"swagger/handler"
)

//	@title			example document
//	@version		1.0
//	@description	test of swag

// @host		127.0.0.1:8080
// @BasePath	/api/v1
func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", handler.Helloworld)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	_ = r.Run("127.0.0.1:8080")

}
