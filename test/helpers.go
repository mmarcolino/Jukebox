package test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func NewMockDB(t *testing.T) (sqlmock.Sqlmock, *sqlx.DB) {
	t.Helper()

	db, mock, err := sqlmock.New()
	require.Nil(t, err)

	return mock, sqlx.NewDb(db, "mock")
}
