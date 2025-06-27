-- name: GetTracksByIDs :many
SELECT id, title, artist, album, genre, duration
FROM public.tracks
WHERE id = ANY($1::text[]);
