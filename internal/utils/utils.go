package utils

import "github.com/marcolino/jukebox/gen/openapi"

func ToOptString (s string) openapi.OptString{
	if s == ""{
		return openapi.OptString{}
	}
	return openapi.NewOptString(s)
}