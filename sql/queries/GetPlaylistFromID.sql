-- name: GetPlaylistFromID :one
SELECT * FROM public.playlist WHERE id = $1;