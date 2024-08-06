-- name: GetRoom :one
SELECT id, theme FROM rooms where id = $1;