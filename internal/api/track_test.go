package api_test

import (
	"context"
	"testing"

	"github.com/marcolino/jukebox/gen/openapi"
	"github.com/marcolino/jukebox/internal/api"
	"github.com/marcolino/jukebox/internal/domain/entity"
	"github.com/marcolino/jukebox/internal/utils"
	"github.com/marcolino/jukebox/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetTracks(t *testing.T) {
	ctx := context.Background()

	tracksResponse := []entity.Track{
		{
			Title:    "Next Semester",
			Artist:   "Twenty One Pilots",
			Album:    "Clancy",
			Genre:    "Rock",
			Duration: 249,
		},
		{
			Title:    "todo dia",
			Artist:   "terraplana",
			Album:    "natural",
			Genre:    "shoegaze",
			Duration: 229,
		},
	}

	var expectedTracks openapi.GetTracksOKApplicationJSON = []openapi.Track{
		{
			Title:    "Next Semester",
			Artist:   "Twenty One Pilots",
			Album:    utils.ToOptString("Clancy"),
			Genre:    utils.ToOptString("Rock"),
			Duration: 249,
		},
		{
			Title:    "todo dia",
			Artist:   "terraplana",
			Album:    utils.ToOptString("natural"),
			Genre:    utils.ToOptString("shoegaze"),
			Duration: 229,
		},
	}

	mockTracksHandler := mocks.NewMockTracks(t)
	mockTracksHandler.On("GetTracks", ctx).Return(tracksResponse, nil)

	queue := mocks.NewMockQueue(t)
	handler := api.NewHandler(mockTracksHandler, mocks.NewMockPlaylists(t), queue)

	tracks, err := handler.GetTracks(ctx)
	assert.NoError(t, err)
	assert.Equal(t, &expectedTracks, tracks)
}

func TestPostTrack(t *testing.T) {
	ctx := context.Background()

	req := &openapi.PostTracksReq{
		Title:    "Next Semester",
		Artist:   "Twenty One Pilots",
		Album:    utils.ToOptString("Clancy"),
		Genre:    utils.ToOptString("Rock"),
		Duration: 249,
	}

	expectedTrack := entity.Track{
		Title:    "Next Semester",
		Artist:   "Twenty One Pilots",
		Album:    "Clancy",
		Genre:    "Rock",
		Duration: 249,
	}

	successRes := &openapi.PostTracksCreated{}

	mockTracksHandler := mocks.NewMockTracks(t)
	mockTracksHandler.On("PostTrack", ctx, expectedTrack).Return(nil)

	queue := mocks.NewMockQueue(t)
	handler := api.NewHandler(mockTracksHandler, mocks.NewMockPlaylists(t), queue)

	res, err := handler.PostTracks(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, successRes, res)
}
