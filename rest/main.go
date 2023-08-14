package main

import (
	"rest/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	albumCtrl := controller.NewAlbumsController()

	router.GET("/albums", albumCtrl.GetAlbums)
	router.GET("/albums/:id", albumCtrl.GetAlbumByID)
	router.POST("/albums", albumCtrl.PostAlbums)

	_ = router.Run("localhost:8080")
}
