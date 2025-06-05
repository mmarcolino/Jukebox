package utils

import (
	"database/sql"

	"github.com/marcolino/jukebox/gen/openapi"
)

func ToOptString(s string) openapi.OptString {
	if s == "" {
		return openapi.OptString{}
	}
	return openapi.NewOptString(s)
}

func FromNullStr(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}

	return ""
}
