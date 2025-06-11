package api

import (
	"context"

	"github.com/marcolino/jukebox/gen/openapi"
	"github.com/marcolino/jukebox/internal/domain/entity"
	"github.com/marcolino/jukebox/internal/domain/gateway"
)

type Handler struct {
	tracksHandler gateway.Tracks
}

func NewHandler(module gateway.Tracks) *Handler {
	return &Handler{tracksHandler: module}
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
