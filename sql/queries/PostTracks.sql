-- name: PostTracks :exec
INSERT INTO public.tracks (id, title, artist, album, duration, genre) VALUES($1, $2, $3, $4, $5, $6);