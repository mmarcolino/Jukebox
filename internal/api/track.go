package api

import (
	"context"

	"github.com/marcolino/jukebox/gen/openapi"
	"github.com/marcolino/jukebox/internal/domain/entity"
	"github.com/marcolino/jukebox/internal/utils"
)

func (h *Handler) GetTracks(ctx context.Context) (openapi.GetTracksRes, error) {
	tracks, err := h.tracksHandler.GetTracks(ctx)
	if err != nil {
		return nil, err
	}

	var responseTracks openapi.GetTracksOKApplicationJSON = make([]openapi.Track, len(tracks))

	for i, track := range tracks {
		responseTracks[i] = openapi.Track{
			Artist:   track.Artist,
			Title:    track.Title,
			Album:    utils.ToOptString(track.Album),
			Genre:    utils.ToOptString(track.Genre),
			Duration: track.Duration,
			ID:       track.ID,
		}
	}

	return &responseTracks, nil
}

func (h *Handler) PostTracks(ctx context.Context, req *openapi.PostTracksReq) (openapi.PostTracksRes, error) {
	track := entity.Track{
		Title:    req.Title,
		Artist:   req.Artist,
		Album:    utils.FromOptString(req.Album),
		Genre:    utils.FromOptString(req.Genre),
		Duration: req.Duration,
	}
	err := h.tracksHandler.PostTrack(ctx, track)
	if err != nil {
		return nil, err
	}
	return &openapi.PostTracksCreated{}, nil
}

func (h *Handler) DeleteTrack(ctx context.Context, params openapi.DeleteTrackParams) (openapi.DeleteTrackRes, error) {
	err := h.tracksHandler.DeleteTrack(ctx, entity.Track{ID: params.ID})
	if err != nil {
		return nil, err
	}

	return &openapi.DeleteTrackNoContent{}, nil
}
