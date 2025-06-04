package api

import "github.com/marcolino/jukebox/internal/domain/gateway"

type Handler struct {
	moduleHandler gateway.Tracks
}

func NewHandler(module gateway.Tracks) *Handler {
	return &Handler{moduleHandler: module}
}
