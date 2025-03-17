-- name: CreatePost :one
INSERT INTO posts (title, url, description, published_at, feed_id)
VALUES (
     $1,
     $2,
     $3,
     $4,
     $5
)
RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.*
FROM posts
JOIN feeds ON feeds.id = posts.feed_id
WHERE feeds.user_id  = $1
ORDER BY posts.published_at DESC
LIMIT $2;