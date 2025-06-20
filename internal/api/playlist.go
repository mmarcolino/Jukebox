package api

import (
	"context"

	"github.com/marcolino/jukebox/gen/openapi"
	"github.com/marcolino/jukebox/internal/domain/entity"
)

func (h *Handler) GetPlaylists(ctx context.Context) (openapi.GetPlaylistsRes, error) {
	playlists, err := h.playlistsHandler.GetPlaylists(ctx)
	if err != nil {
		return nil, err
	}

	var responsePlaylists openapi.GetPlaylistsOKApplicationJSON = make([]openapi.Playlist, len(playlists))

	for i, playlist := range playlists {
		responsePlaylists[i] = openapi.Playlist{
			ID:     playlist.ID,
			Name:   playlist.Name,
			Tracks: playlist.Tracks,
		}
	}

	return &responsePlaylists, nil
}

func (h *Handler) PostPlaylist(ctx context.Context, req *openapi.PostPlaylistReq) (openapi.PostPlaylistRes, error) {
	playlist := entity.Playlist{
		Name:   req.Name,
		Tracks: req.Track,
	}

	err := h.playlistsHandler.CreatePlaylist(ctx, playlist)
	if err != nil {
		return nil, err
	}

	return &openapi.PostPlaylistCreated{}, nil
}
