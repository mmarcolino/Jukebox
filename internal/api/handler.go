package api

import (
	"context"

	"github.com/marcolino/jukebox/gen/openapi"
	"github.com/marcolino/jukebox/internal/domain/entity"
	"github.com/marcolino/jukebox/internal/domain/gateway"
)

type Handler struct {
	tracksHandler    gateway.Tracks
	playlistsHandler gateway.Playlists
	queueClient      gateway.Queue
}

func NewHandler(tracks gateway.Tracks, playlists gateway.Playlists, queueClient gateway.Queue) *Handler {
	return &Handler{tracksHandler: tracks, playlistsHandler: playlists, queueClient: queueClient}
}

func (h *Handler) NewError(ctx context.Context, err error) *openapi.ErrorResponseOgenStatusCode {
	statusCode := 500
	errMsg := "INTERNAL_ERROR"

	switch err.Error() {
	case "NOT_FOUND":
		statusCode = 404
		errMsg = entity.ErrNotFound.Error()
	}

	return &openapi.ErrorResponseOgenStatusCode{
		StatusCode: statusCode,
		Response: openapi.ErrorResponseOgen{
			ErrorMessage: errMsg,
			StatusCode:   statusCode,
		},
	}
}
