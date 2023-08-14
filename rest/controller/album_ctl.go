package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest/model"
)

type AlbumsController struct{}

func NewAlbumsController() *AlbumsController {
	return &AlbumsController{}
}

// GetAlbums responds with the list of all albums as JSON.
func (a *AlbumsController) GetAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, model.Albums)
}

// PostAlbums adds an album from JSON received in the request body.
func (a *AlbumsController) PostAlbums(c *gin.Context) {
	var newAlbum model.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	model.Albums = append(model.Albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumsController) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range model.Albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
