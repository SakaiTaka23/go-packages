package controller

import (
	"context"
	"errors"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
	"rest/model"
)

type AlbumsController struct{}

func NewAlbumsController() *AlbumsController {
	return &AlbumsController{}
}

// GetAlbums responds with the list of all albums as JSON.
func (a *AlbumsController) GetAlbums() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, _ struct{}, output *[]model.Album) error {
		*output = model.Albums
		return nil
	})
	u.SetTags("Album")

	return u
}

// PostAlbums adds an album from JSON received in the request body.
func (a *AlbumsController) PostAlbums() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, input model.Album, output *model.Album) error {
		// Check if id is unique.
		for _, album := range model.Albums {
			if album.ID == input.ID {
				return status.AlreadyExists
			}
		}

		// Add the new album to the slice.
		model.Albums = append(model.Albums, input)

		*output = input
		return nil
	})
	u.SetTags("Album")
	u.SetExpectedErrors(status.AlreadyExists)

	return u
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumsController) GetAlbumByID() usecase.Interactor {
	type getAlbumByIDInput struct {
		ID string `path:"id"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, input getAlbumByIDInput, output *model.Album) error {
		for _, album := range model.Albums {
			if album.ID == input.ID {
				*output = album
				return nil
			}
		}
		return status.Wrap(errors.New("album not found"), status.NotFound)
	})
	u.SetTags("Album")
	u.SetExpectedErrors(status.NotFound)

	return u
}
