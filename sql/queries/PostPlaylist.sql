-- name: PostPlaylist :exec
INSERT INTO public.playlist (id, name, tracks) VALUES($1, $2, $3);