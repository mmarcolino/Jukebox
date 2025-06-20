package gateway

import (
	"context"

	"github.com/marcolino/jukebox/internal/domain/entity"
)

type Queue interface {
	AddTrackToQueue(ctx context.Context, track entity.Track) error
	ReceiveTracks(ctx context.Context) ([]entity.Track, error)
}
