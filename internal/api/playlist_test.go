package api_test

import (
	"context"
	"testing"

	"github.com/marcolino/jukebox/gen/openapi"
	"github.com/marcolino/jukebox/internal/api"
	"github.com/marcolino/jukebox/internal/domain/entity"
	"github.com/marcolino/jukebox/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetPlaylist(t *testing.T){
	ctx := context.Background()

	getPlaylistsResponse := []entity.Playlist{
		{
			ID: "01JX3872K622GTRCCVXHXVP8ZY",
			Name: "Playlist1",
			Tracks: []string {"T1", "T2"},
		},
		{
			ID: "01JX3872K622GTRCCVXHXVP8ZZ",
			Name: "Playlist2",
			Tracks: []string {"T3", "T24"},
		},
	}
	var expectedPlaylists openapi.GetPlaylistsOKApplicationJSON = []openapi.Playlist{
		{
			ID: "01JX3872K622GTRCCVXHXVP8ZY",
			Name: "Playlist1",
			Tracks: []string {"T1", "T2"},
		},
		{
			ID: "01JX3872K622GTRCCVXHXVP8ZZ",
			Name: "Playlist2",
			Tracks: []string {"T3", "T24"},
		},
	}
	for name, scenario := range map[string]struct{
		returnData       []entity.Playlist
		expectedResponse openapi.GetPlaylistsRes
		expectedError    error
	}{
		"success":{
			returnData: getPlaylistsResponse,
			expectedResponse: &expectedPlaylists,
			expectedError: nil,
		},
		"failure":{
			returnData: nil,
			expectedResponse: nil,
			expectedError: entity.GenericErr,
		},
	}{
		t.Run(name, func(t *testing.T) {
			mockPlaylists := mocks.NewMockPlaylists(t)
			mockPlaylists.On("GetPlaylists", ctx).Return(scenario.returnData, scenario.expectedError)
	
			handler := api.NewHandler(mocks.NewMockTracks(t), mockPlaylists, mocks.NewMockQueue(t))
			response, err := handler.GetPlaylists(ctx)
	
			assert.Equal(t, scenario.expectedResponse, response)
			assert.ErrorIs(t, err, scenario.expectedError)
		})
	}
}

func TestPostPlaylist(t *testing.T){
	ctx := context.Background()

	for name, scenario := range map[string]struct{
		expectedPlaylist entity.Playlist
		postPlaylistParam openapi.PostPlaylistReq
		expectedReturn openapi.PostPlaylistRes
		expectedError error
	}{
		"success":{
			expectedPlaylist: entity.Playlist{
				Name: "TestPlaylist",
				Tracks: []string{"track1", "track2"},
			},
			postPlaylistParam: openapi.PostPlaylistReq{
				Name: "TestPlaylist",
				Track: []string{"track1", "track2"},
			},
			expectedReturn: &openapi.PostPlaylistCreated{},
			expectedError: nil,
		},
		"failure":{
			expectedPlaylist: entity.Playlist{},
			postPlaylistParam: openapi.PostPlaylistReq{},
			expectedReturn: nil,
			expectedError: entity.GenericErr,
		},
	}{
		t.Run(name, func(t *testing.T) {
			mockPlaylists := mocks.NewMockPlaylists(t)
			mockPlaylists.On("CreatePlaylist", ctx, scenario.expectedPlaylist).Return(scenario.expectedError)

			handler := api.NewHandler(mocks.NewMockTracks(t), mockPlaylists, mocks.NewMockQueue(t))
			response, err := handler.PostPlaylist(ctx, &scenario.postPlaylistParam)
			assert.Equal(t, scenario.expectedReturn, response)
			assert.ErrorIs(t, err, scenario.expectedError)
		})
	}
}

func TestExecutePlaylist(t *testing.T){
	ctx := context.Background()
	
	mockPlaylist := entity.Playlist{
		ID: "01JX3872K622GTRCCVXHXVP8ZY",
		Name: "TestPlaylist",
		Tracks: []string{"01JX3872K622GTRCCVXHXVP8ZX", "01JX3872K622GTRCCVXHXVP8ZZ"},
	}
	mockTracks := []entity.Track{
		{
			ID: "01JX3872K622GTRCCVXHXVP8ZX",
			Title: "Track1",
			Artist: "ART1",
			Album: "ALB1",
			Genre: "G1",
			Duration: 111,
		},
		{
			ID: "01JX3872K622GTRCCVXHXVP8ZZ",
			Title: "Track2",
			Artist: "ART2",
			Album: "ALB2",
			Genre: "G2",
			Duration: 222,
		},
	}

	for name, scenario := range map[string]struct{
		expectedID openapi.ExecutePlaylistParams
		getPlaylistResponse entity.Playlist
		getTracksParam []string
		getTracksResponse []entity.Track
		addToQueueParam []entity.Track
		getPlaylistErr error
		getTracksErr error
		addToQueueErr error
		expectedErr error
		expectedResponse openapi.ExecutePlaylistRes
	}{
		"success":{
			expectedID: openapi.ExecutePlaylistParams{ID: mockPlaylist.ID},
			getPlaylistResponse: mockPlaylist,
			getTracksParam: mockPlaylist.Tracks,
			getTracksResponse: mockTracks,
			addToQueueParam: mockTracks,
			getPlaylistErr: nil,
			getTracksErr: nil,
			addToQueueErr: nil,
			expectedErr: nil,
			expectedResponse: &openapi.ExecutePlaylistOK{},
		},
		"get_playlist_failure":{
			expectedID: openapi.ExecutePlaylistParams{},
			getPlaylistResponse: entity.Playlist{},
			getTracksParam: mockPlaylist.Tracks,
			getTracksResponse: mockTracks,
			getPlaylistErr: entity.GenericErr,
			getTracksErr: nil,
			addToQueueErr: nil,
			expectedErr: entity.GenericErr,
			expectedResponse: nil,
		},
		"get_tracks_failure":{
			expectedID: openapi.ExecutePlaylistParams{ID: mockPlaylist.ID},
			getPlaylistResponse: mockPlaylist,
			getTracksParam: mockPlaylist.Tracks,
			getTracksResponse: []entity.Track{},
			getPlaylistErr: nil,
			getTracksErr: entity.GenericErr,
			addToQueueErr: nil,
			expectedErr: entity.GenericErr,
			expectedResponse: nil,
		},
		"add_to_queue_failure":{
			expectedID: openapi.ExecutePlaylistParams{ID: mockPlaylist.ID},
			getPlaylistResponse: mockPlaylist,
			getTracksParam: mockPlaylist.Tracks,
			getTracksResponse: mockTracks,
			addToQueueParam: mockTracks,
			getPlaylistErr: nil,
			getTracksErr: nil,
			addToQueueErr: entity.GenericErr,
			expectedErr: entity.GenericErr,
			expectedResponse: nil,
		},
	}{
		t.Run(name, func(t *testing.T) {
			mockPlaylistHandler := mocks.NewMockPlaylists(t)
			mockTracksHandler := mocks.NewMockTracks(t)
			mockQueue := mocks.NewMockQueue(t)

			mockPlaylistHandler.On("GetPlaylistFromID", ctx, scenario.expectedID.ID).Return(scenario.getPlaylistResponse, scenario.getPlaylistErr)
			if scenario.getPlaylistErr == nil {
					mockTracksHandler.On("GetTracksFromPlaylist", ctx, scenario.getTracksParam).Return(scenario.getTracksResponse, scenario.getTracksErr)	
					if len(scenario.addToQueueParam) > 0 {
					if scenario.getTracksErr == nil{
						for _, track := range scenario.addToQueueParam{
							mockQueue.On("AddTrackToQueue", ctx, track).Return(scenario.addToQueueErr)
							if scenario.addToQueueErr != nil{
								break
							}	
						}
					}
				}
			}
			
			handler := api.NewHandler(mockTracksHandler, mockPlaylistHandler, mockQueue)
			response, err := handler.ExecutePlaylist(ctx, scenario.expectedID)
			assert.Equal(t, scenario.expectedResponse, response)
			assert.ErrorIs(t, err, scenario.expectedErr)
		})
	}
}