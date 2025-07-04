package postgres

import (
	"context"

	"github.com/marcolino/jukebox/gen/sqlc"
	"github.com/marcolino/jukebox/internal/domain/entity"
	"github.com/marcolino/jukebox/internal/utils"
	"github.com/oklog/ulid/v2"
)

func (h *PostgresHandler) GetTracks(ctx context.Context) ([]entity.Track, error) {
	persistedTracks, err := h.queries.GetTracks(ctx)
	if err != nil {
		return nil, err
	}

	if len(persistedTracks) <= 0 {
		return nil, entity.ErrNotFound
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

func (h *PostgresHandler) PostTrack(ctx context.Context, track entity.Track) error {
	return h.queries.PostTracks(ctx, sqlc.PostTracksParams{
		ID:       ulid.Make().String(),
		Title:    track.Title,
		Artist:   track.Artist,
		Duration: int32(track.Duration),
		Album:    utils.ToNullString(track.Album),
		Genre:    utils.ToNullString(track.Genre),
	})
}

func (h *PostgresHandler) DeleteTrack(ctx context.Context, track entity.Track) error {
	return h.queries.DeleteTrack(ctx, track.ID)
}

func (h *PostgresHandler) UpdateTrack(ctx context.Context, track entity.Track) error {
	return h.queries.UpdateTrack(ctx, sqlc.UpdateTrackParams{
		ID:       track.ID,
		Title:    track.Title,
		Artist:   track.Artist,
		Duration: int32(track.Duration),
		Album:    utils.ToNullString(track.Album),
		Genre:    utils.ToNullString(track.Genre),
	})
}

func (h *PostgresHandler) GetTracksFromPlaylist(ctx context.Context, tracksIDs []string) ([]entity.Track, error) {
	persistedTracks, err := h.queries.GetTracksByIDs(ctx, tracksIDs)
	if err != nil {
		return nil, err
	}

	if len(persistedTracks) <= 0 {
		return nil, entity.ErrNotFound
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
