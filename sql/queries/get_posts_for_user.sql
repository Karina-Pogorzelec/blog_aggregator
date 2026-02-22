-- name: GetPostsForUser :many
select 
    posts.*
from posts
join feed_follows on posts.feed_id = feed_follows.feed_id
join users on feed_follows.user_id = users.id
where users.id = $1
order by posts.created_at desc
limit $2;