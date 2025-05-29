package api

import (
	"context"

	"github.com/marcolino/jukebox/gen/openapi"
	"github.com/marcolino/jukebox/internal/domain/entity"
	"github.com/marcolino/jukebox/internal/utils"
) 


func (h *Handler) GetTracks(ctx context.Context) ([]openapi.Track, error){
	tracks := []entity.Track{{
		Artist: "Terraplana",
		Title: "Conversas",
		Album: "Olhar pra trás",
		Genre: "Shoegaze",
		Duration: 180,
	},
	{
		Artist: "Terraplana",
		Title: "Conversas",
		Album: "Olhar pra trás",
		Genre: "Shoegaze",
		Duration: 180,
	}}
	var responseTracks []openapi.Track = make([]openapi.Track,len(tracks))
	for i, track := range tracks{
		responseTracks[i] = openapi.Track{
			Artist: track.Artist,
			Title: track.Title,
			Album: utils.ToOptString(track.Album),
			Genre: utils.ToOptString(track.Genre),
			Duration: track.Duration,
		}
	}
	return responseTracks,nil
}