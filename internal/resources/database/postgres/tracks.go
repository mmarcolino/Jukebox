package postgres

import (
	"context"

	"github.com/marcolino/jukebox/internal/domain/entity"
	"github.com/marcolino/jukebox/internal/utils"
)

func (h *PostgresHandler) GetTracks(ctx context.Context) ([]entity.Track, error) {
	persistedTracks, err := h.queries.GetTracks(ctx)
	if err != nil {
		return nil, err
	}

	var tracks []entity.Track = make([]entity.Track, len(persistedTracks))

	for i, track := range persistedTracks {
		tracks[i] = entity.Track{
			ID:       track.ID,
			Title:    track.Title,
			Artist:   track.Artist,
			Album:    utils.FromNullStr(track.Album),
			Genre:    utils.FromNullStr(track.Genre),
			Duration: int(track.Duration),
		}
	}

	return tracks, nil
}
