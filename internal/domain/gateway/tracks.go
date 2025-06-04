package gateway

import (
	"context"

	"github.com/marcolino/jukebox/internal/domain/entity"
)

type Tracks interface {
	GetTracks(context.Context) ([]entity.Track, error)
}
