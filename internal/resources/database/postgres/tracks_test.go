package postgres_test

import (
	"context"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcolino/jukebox/internal/domain/entity"
	"github.com/marcolino/jukebox/internal/resources/database/postgres"
	"github.com/marcolino/jukebox/test"
	"github.com/stretchr/testify/assert"
)

func TestGetTracks(t *testing.T) {
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

	columns := []string{"id", "title", "artist", "album", "duration", "genre"}

	rowsData := [][]driver.Value{
		{
			"01JX3872K622GTRCCVXHXVP8ZY", "Next Semester", "Twenty One Pilots", "Clancy", int32(249), "Rock",
		},
		{
			"01JX38A0KPBN0RTDEWRMYFN0K2", "todo dia", "terraplana", "natural", int32(229), "shoegaze",
		},
	}

	mock, db := test.NewMockDB(t)

	expectedRows := sqlmock.NewRows(columns)

	for _, row := range rowsData {
		expectedRows.AddRow(row...)
	}

	mock.ExpectQuery("-- name: GetTracks :many").WillReturnRows(expectedRows)

	postgresHandler := postgres.New(db)

	tracks, err := postgresHandler.GetTracks(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedTracks, tracks)
}

func TestPostTrack(t *testing.T) {
	ctx := context.Background()

	data := entity.Track{
		ID:       "01JX3872K622GTRCCVXHXVP8ZY",
		Title:    "Next Semester",
		Artist:   "Twenty One Pilots",
		Album:    "Clancy",
		Genre:    "Rock",
		Duration: 249,
	}

	mock, db := test.NewMockDB(t)
	mock.ExpectExec("^-- name: PostTracks :exec INSERT INTO public.tracks*").
		WithArgs(sqlmock.AnyArg(), data.Title, data.Artist, data.Album, data.Duration, data.Genre).
		WillReturnResult(driver.ResultNoRows)

	postgresHandler := postgres.New(db)

	err := postgresHandler.PostTrack(ctx, data)
	assert.NoError(t, err)

}
