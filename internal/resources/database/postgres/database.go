package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/marcolino/jukebox/gen/sqlc"
	"github.com/marcolino/jukebox/internal/domain/gateway"
)

type PostgresHandler struct {
	db      *sqlx.DB
	queries *sqlc.Queries
}

var _ gateway.Tracks = (*PostgresHandler)(nil)
var _ gateway.Playlists = (*PostgresHandler)(nil)

func New(db *sqlx.DB) *PostgresHandler {
	return &PostgresHandler{db: db, queries: sqlc.New(db)}
}
