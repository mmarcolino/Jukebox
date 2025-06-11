package entity

import "errors"

var (
	ErrNotFound = errors.New("NOT_FOUND")
	GenericErr  = errors.New("GENERIC_ERROR")
)
