-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (user_id, feed_id)
    VALUES ($1, $2)
    RETURNING *
)

SELECT 
  inserted_feed_follow.*,
  feeds.name AS feed_name,
  users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users ON user_id = users.id
INNER JOIN feeds ON feed_id = feeds.id;



-- name: GetFeedFollowsForUser :many
SELECT 
  feed_follows.*,
  feeds.name AS feed_name,
  users.name AS user_name
FROM feed_follows
INNER JOIN users ON user_id = users.id
INNER JOIN feeds ON feed_id = feeds.id
WHERE feed_follows.user_id = $1;


-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE feed_follows.user_id = $1 AND feed_follows.feed_id = (SELECT id
FROM feeds
WHERE feeds.url = $2);
