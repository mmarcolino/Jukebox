-- name: DeleteTrack :exec
DELETE FROM public.tracks WHERE id = $1;