package api

import (
	"context"

	"github.com/marcolino/jukebox/gen/openapi"
	"github.com/marcolino/jukebox/internal/domain/entity"
	"github.com/marcolino/jukebox/internal/utils"
)

func (h *Handler) GetTracks(ctx context.Context) ([]openapi.Track, error) {
	tracks, err := h.moduleHandler.GetTracks(ctx)
	if err != nil {
		return nil, err
	}

	var responseTracks []openapi.Track = make([]openapi.Track, len(tracks))

	for i, track := range tracks {
		responseTracks[i] = openapi.Track{
			Artist:   track.Artist,
			Title:    track.Title,
			Album:    utils.ToOptString(track.Album),
			Genre:    utils.ToOptString(track.Genre),
			Duration: track.Duration,
		}
	}
	return responseTracks, nil
}

func (h *Handler) PostTracks(ctx context.Context, req *openapi.PostTracksReq) (openapi.PostTracksRes, error){
	track := entity.Track{
		Title: req.Title,
		Artist: req.Artist,
		Album: utils.FromOptString(req.Album),
		Genre: utils.FromOptString(req.Genre),
		Duration: req.Duration,
	}
	err := h.moduleHandler.PostTrack(ctx, track)
	if err != nil{
		return nil, err
	}
	return &openapi.PostTracksCreated{}, nil
}
