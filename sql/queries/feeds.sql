-- name: CreateFeed :one
INSERT INTO feeds (id, name, creted_at, updated_at, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE id = $1;