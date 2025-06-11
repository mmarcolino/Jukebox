package postgres_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcolino/jukebox/gen/sqlc"
	"github.com/marcolino/jukebox/internal/domain/entity"
	"github.com/marcolino/jukebox/internal/resources/database/postgres"
	"github.com/marcolino/jukebox/test"
	"github.com/stretchr/testify/assert"
)

func TestGetTracks(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	expectedTracks := []entity.Track{
		{
			ID:       "01JX3872K622GTRCCVXHXVP8ZY",
			Title:    "Next Semester",
			Artist:   "Twenty One Pilots",
			Album:    "Clancy",
			Genre:    "Rock",
			Duration: 249,
		},
		{
			ID:       "01JX38A0KPBN0RTDEWRMYFN0K2",
			Title:    "todo dia",
			Artist:   "terraplana",
			Album:    "natural",
			Genre:    "shoegaze",
			Duration: 229,
		},
	}

	successRowsData := [][]driver.Value{
		{
			"01JX3872K622GTRCCVXHXVP8ZY", "Next Semester", "Twenty One Pilots", "Clancy", int32(249), "Rock",
		},
		{
			"01JX38A0KPBN0RTDEWRMYFN0K2", "todo dia", "terraplana", "natural", int32(229), "shoegaze",
		},
	}

	columns := []string{"id", "title", "artist", "album", "duration", "genre"}

	for name, scenario := range map[string]struct {
		expepectedRowsData [][]driver.Value
		expectedData       []entity.Track
		queryError         error
		expectedError      error
	}{
		"success": {
			expepectedRowsData: successRowsData,
			expectedData:       expectedTracks,
			queryError:         nil,
			expectedError:      nil,
		},
		"not-found": {
			expepectedRowsData: [][]driver.Value{},
			expectedData:       nil,
			queryError:         nil,
			expectedError:      entity.ErrNotFound,
		},
		"generic-error": {
			expepectedRowsData: [][]driver.Value{},
			expectedData:       nil,
			queryError:         entity.GenericErr,
			expectedError:      entity.GenericErr,
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			mock, db := test.NewMockDB(t)
			expectedRows := sqlmock.NewRows(columns)

			for _, row := range scenario.expepectedRowsData {
				expectedRows.AddRow(row...)
			}

			mock.ExpectQuery("-- name: GetTracks :many").WillReturnRows(expectedRows).WillReturnError(scenario.queryError)

			postgresHandler := postgres.New(db)
			tracks, err := postgresHandler.GetTracks(ctx)

			assert.ErrorIs(t, err, scenario.expectedError)
			assert.Equal(t, scenario.expectedData, tracks)
		})
	}

}

func TestPostTrack(t *testing.T) {
	ctx := context.Background()

	for name, scenario := range map[string]struct {
		data          entity.Track
		sqlcData      sqlc.PostTracksParams
		queryError    error
		expectedError error
	}{
		"success": {
			data: entity.Track{
				ID:       "01JX3872K622GTRCCVXHXVP8ZY",
				Title:    "Next Semester",
				Artist:   "Twenty One Pilots",
				Album:    "Clancy",
				Genre:    "Rock",
				Duration: 249,
			},
			sqlcData: sqlc.PostTracksParams{
				ID:       "01JX3872K622GTRCCVXHXVP8ZY",
				Title:    "Next Semester",
				Artist:   "Twenty One Pilots",
				Album:    sql.NullString{String: "Clancy", Valid: true},
				Genre:    sql.NullString{String: "Rock", Valid: true},
				Duration: 249,
			},
			queryError:    nil,
			expectedError: nil,
		},
		"generic-error": {
			data:          entity.Track{},
			sqlcData:      sqlc.PostTracksParams{},
			queryError:    entity.GenericErr,
			expectedError: entity.GenericErr,
		},
	} {
		t.Run(name, func(t *testing.T) {
			mock, db := test.NewMockDB(t)
			mock.ExpectExec("^-- name: PostTracks :exec INSERT INTO public.tracks*").
				WithArgs(sqlmock.AnyArg(), scenario.sqlcData.Title, scenario.sqlcData.Artist, sqlmock.AnyArg(), scenario.sqlcData.Duration, sqlmock.AnyArg()).
				WillReturnResult(driver.ResultNoRows).WillReturnError(scenario.queryError)

			postgresHandler := postgres.New(db)
			err := postgresHandler.PostTrack(ctx, scenario.data)
			assert.ErrorIs(t, err, scenario.expectedError)
		})
	}

}

func TestDeleteTrack(t *testing.T) {
	ctx := context.Background()

	columns := []string{"id", "title", "artist", "album", "duration", "genre"}

	successRowsData := [][]driver.Value{
		{
			"01JX3872K622GTRCCVXHXVP8ZY", "Next Semester", "Twenty One Pilots", "Clancy", int32(249), "Rock",
		},
	}

	for name, scenario := range map[string]struct {
		rows          [][]driver.Value
		data          entity.Track
		queryError    error
		expectedError error
	}{
		"sucess": {
			rows: successRowsData,
			data: entity.Track{
				ID:       "01JX3872K622GTRCCVXHXVP8ZY",
				Title:    "Next Semester",
				Artist:   "Twenty One Pilots",
				Album:    "Clancy",
				Genre:    "Rock",
				Duration: 249,
			},
			queryError:    nil,
			expectedError: nil,
		},

		"generic-error": {
			rows:          [][]driver.Value{},
			data:          entity.Track{},
			queryError:    entity.GenericErr,
			expectedError: entity.GenericErr,
		},
	} {
		t.Run(name, func(t *testing.T) {

			mock, db := test.NewMockDB(t)

			expectedRows := sqlmock.NewRows(columns)

			for _, row := range scenario.rows {
				expectedRows.AddRow(row...)
			}

			mock.ExpectExec("^-- name: DeleteTrack :exec DELETE FROM public.tracks*").
				WithArgs(scenario.data.ID).
				WillReturnResult(driver.ResultNoRows).
				WillReturnError(scenario.queryError)

			postgresHandler := postgres.New(db)

			err := postgresHandler.DeleteTrack(ctx, scenario.data)
			assert.ErrorIs(t, err, scenario.expectedError)
		})
	}

}
