package main

import (
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/swgui/v4emb"
	"log"
	"net/http"
	"rest/controller"

	"github.com/swaggest/rest/web"
)

func main() {
	// Service initializes router with required middlewares.
	service := web.NewService(openapi3.NewReflector())
	// It allows OpenAPI configuration.
	service.OpenAPISchema().SetTitle("Albums API")
	service.OpenAPISchema().SetDescription("This service provides API to manage albums.")
	service.OpenAPISchema().SetVersion("v1.0.0")

	albumCtrl := controller.NewAlbumsController()

	service.Get("/albums", albumCtrl.GetAlbums())
	service.Post("/albums", albumCtrl.PostAlbums(), nethttp.SuccessStatus(http.StatusCreated))
	service.Get("/albums/{id}", albumCtrl.GetAlbumByID())

	service.Docs("/docs", v4emb.New)

	log.Println("Starting service")
	if err := http.ListenAndServe("localhost:8080", service); err != nil {
		log.Fatal(err)
	}
}
