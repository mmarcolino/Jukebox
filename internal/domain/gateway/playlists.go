package gateway

import (
	"context"

	"github.com/marcolino/jukebox/internal/domain/entity"
)

type Playlists interface {
	GetPlaylists(context.Context) ([]entity.Playlist, error)
	CreatePlaylist(context.Context, entity.Playlist) error
}
