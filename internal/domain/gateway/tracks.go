package gateway

import (
	"context"

	"github.com/marcolino/jukebox/internal/domain/entity"
)

type Tracks interface {
	GetTracks(context.Context) ([]entity.Track, error)
	PostTrack(context.Context, entity.Track) error
	DeleteTrack(context.Context, entity.Track) error
	UpdateTrack(context.Context, entity.Track) error
}
