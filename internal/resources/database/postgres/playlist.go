package postgres

import (
	"context"

	"github.com/marcolino/jukebox/gen/sqlc"
	"github.com/marcolino/jukebox/internal/domain/entity"
	"github.com/oklog/ulid/v2"
)

func (p *PostgresHandler) GetPlaylists(ctx context.Context) ([]entity.Playlist, error) {
	persistedPlaylists, err := p.queries.GetPlaylist(ctx)
	if err != nil {
		return nil, err
	}

	if len(persistedPlaylists) <= 0 {
		return nil, entity.ErrNotFound
	}

	var playlists []entity.Playlist = make([]entity.Playlist, len(persistedPlaylists))
	for i, persistedPlaylist := range persistedPlaylists {
		playlists[i] = entity.Playlist{
			ID:     persistedPlaylist.ID,
			Name:   persistedPlaylist.Name,
			Tracks: persistedPlaylist.Tracks,
		}
	}

	return playlists, nil
}

func (p *PostgresHandler) CreatePlaylist(ctx context.Context, playlist entity.Playlist) error {
	return p.queries.PostPlaylist(ctx, sqlc.PostPlaylistParams{
		ID:     ulid.Make().String(),
		Name:   playlist.Name,
		Tracks: playlist.Tracks,
	})
}

func (p *PostgresHandler) GetPlaylistFromID(ctx context.Context, id string) (entity.Playlist, error) {
	persistedPlaylist, err := p.queries.GetPlaylistFromID(ctx, id)
	if err != nil {
		return entity.Playlist{}, err
	}

	return entity.Playlist{
		ID:     persistedPlaylist.ID,
		Name:   persistedPlaylist.Name,
		Tracks: persistedPlaylist.Tracks,
	}, err

}
