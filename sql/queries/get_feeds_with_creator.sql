-- name: GetFeedsWithCreator :many
SELECT
  feeds.name,
  feeds.url,
  users.name AS username
FROM feeds
JOIN users ON users.id = feeds.user_id
ORDER BY feeds.created_at;