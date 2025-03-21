-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
     gen_random_uuid(),
     NOW(),
     NOW(),
     $1,
     $2,
     $3
)
RETURNING *;

-- name: GetFeedsWithUser :many
SELECT 
  feeds.name, 
  feeds.url,
  users.name AS username
FROM feeds
JOIN users ON feeds.user_id = users.id;


-- name: GetFeedByUrl :one
SELECT *
FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds 
SET last_fetched_at = NOW(), updated_at = NOW()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * 
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST 
LIMIT 1;
