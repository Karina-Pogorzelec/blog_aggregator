-- name: GetFeedByURL :one
SELECT * FROM feeds where url = $1;