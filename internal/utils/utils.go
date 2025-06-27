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

func FromOptString(o openapi.OptString) string {
	if o.IsSet() {
		return o.Value
	}
	return ""
}

func FromNullStr(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}

	return ""
}

func ToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{Valid: true, String: s}
}

func PointerToString(p *string) string {
	if p == nil {
		return ""
	}

	return *p
}

func StringToPointer(s string) *string{
	if s == ""{
		return nil
	}
	return &s
}
