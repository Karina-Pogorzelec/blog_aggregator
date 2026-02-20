-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
select 
    inserted_feed_follow.*, 
    users.name as username, 
    feeds.name as feedname
from inserted_feed_follow
join users on users.id = inserted_feed_follow.user_id
join feeds on feeds.id = inserted_feed_follow.feed_id;