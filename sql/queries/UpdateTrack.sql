-- name: UpdateTrack :exec
UPDATE public.tracks 
SET title=$2, artist=$3, album=$4, duration=$5, genre=$6
WHERE id=$1;