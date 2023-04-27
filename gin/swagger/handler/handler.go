package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Helloworld
//
//	@Summary		return helloworld
//	@Schemes		http
//	@Description	do hello world
//	@Tags			greeting
//	@Produce		json
//	@Success		200	{string}	helloworld
//	@Router			/example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
