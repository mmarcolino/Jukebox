package api

import "github.com/marcolino/jukebox/internal/domain/gateway"

type Handler struct {
	tracksHandler gateway.Tracks
}

func NewHandler(module gateway.Tracks) *Handler {
	return &Handler{tracksHandler: module}
}
