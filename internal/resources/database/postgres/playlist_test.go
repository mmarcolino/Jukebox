package postgres_test

import (
	"context"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcolino/jukebox/gen/sqlc"
	"github.com/marcolino/jukebox/internal/domain/entity"
	"github.com/marcolino/jukebox/internal/resources/database/postgres"
	"github.com/marcolino/jukebox/test"
	"github.com/stretchr/testify/assert"
)

func TestGetPlaylists(t *testing.T){
	ctx := context.Background()
	expectedPlaylists := []entity.Playlist{
		{
			ID:  "01JX3872K622GTRCCVXHXVP8ZY",
			Name: "Playlist1",
			Tracks: []string{"track1", "track2", "track3"},
		},
		{
			ID:  "01JX38A0KPBN0RTDEWRMYFN0K2",
			Name: "Playlist2",
			Tracks: []string{"track4", "track5", "track6"},
		},
	}

	successRowsData := [][]driver.Value{
		{
			"01JX3872K622GTRCCVXHXVP8ZY", "Playlist1", "{track1,track2,track3}",
		},
		{
			"01JX38A0KPBN0RTDEWRMYFN0K2", "Playlist2", "{track4,track5,track6}",
		},
	}
	columns := []string{"id", "name", "tracks"}

	for name, scenario := range map[string]struct{
		expectedRowsData [][]driver.Value
		expectedData     []entity.Playlist
		queryError       error
		expectedError    error 
	}{
		"sucess":{
			expectedRowsData: successRowsData,
			expectedData:     expectedPlaylists,
			queryError:       nil,
			expectedError:    nil,
		},
		"not-found":{
			expectedRowsData: [][]driver.Value{},
			expectedData:     nil,
			queryError:       nil,
			expectedError:    entity.ErrNotFound,
		},
		"generic-error":{
			expectedRowsData: [][]driver.Value{},
			expectedData:     nil,
			queryError:       entity.GenericErr,
			expectedError:    entity.GenericErr,
		},
	}{
		t.Run(name, func(t *testing.T){
			mock, db := test.NewMockDB(t)
			expectedRows := sqlmock.NewRows(columns)

			for _, row := range scenario.expectedRowsData{
				expectedRows.AddRow(row...)
			}

			mock.ExpectQuery("-- name: GetPlaylist :many").WillReturnRows(expectedRows).WillReturnError(scenario.queryError)
			
			postgresHandler := postgres.New(db)
			playlists, err := postgresHandler.GetPlaylists(ctx)

			assert.ErrorIs(t, err, scenario.expectedError)
			assert.Equal(t, scenario.expectedData, playlists)
		
		})
		
	}
}

func TestCreatePlaylist(t *testing.T){
	ctx := context.Background()

	for name, scenario := range map[string]struct{
		data          entity.Playlist
		sqlcData      sqlc.PostPlaylistParams
		queryError    error
		expectedError error
	}{
		"success": {
			data: entity.Playlist{
				ID:  "01JX3872K622GTRCCVXHXVP8ZY",
				Name: "Playlist1",
				Tracks: []string{"track1", "track2", "track3"},
			},
			sqlcData: sqlc.PostPlaylistParams{
				ID:  "01JX3872K622GTRCCVXHXVP8ZY",
				Name: "Playlist1",
				Tracks: []string{"track1", "track2", "track3"},
			},
			queryError:    nil,
			expectedError: nil,
		},
		"generic-error": {
			data:          entity.Playlist{},
			sqlcData:      sqlc.PostPlaylistParams{},
			queryError:    entity.GenericErr,
			expectedError: entity.GenericErr,
		},
	}{
		t.Run(name, func(t *testing.T) {
			mock, db := test.NewMockDB(t)
			mock.ExpectExec("^-- name: PostPlaylist :exec INSERT INTO public.playlist*").
				WithArgs(sqlmock.AnyArg(), scenario.sqlcData.Name, sqlmock.AnyArg()).
				WillReturnResult(driver.ResultNoRows).WillReturnError(scenario.queryError)

			postgresHandler := postgres.New(db)
			err := postgresHandler.CreatePlaylist(ctx, scenario.data)
			assert.ErrorIs(t, err, scenario.expectedError)
		})
	}
}