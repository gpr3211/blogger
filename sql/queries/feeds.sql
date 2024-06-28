-- name: CreateFeed :one

INSERT INTO feeds (id,created_at,updated_at,name,url,user_id,last_fetch)
VALUES($1,$2,$3,$4,$5,$6,$7)
    RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: MakeFetchList :many

SELECT * from feeds
    ORDER BY last_fetch LIMIT $1;


-- name: MarkFeedFetched :one

UPDATE feeds
SET last_fetch = NOW(), updated_at = now() WHERE id = $1
RETURNING *;


/*
-- name: CreateFeed :one
-- name: GetAllFeeds :many
-- name: MakeFetchList :many
-- name: MarkFeedFetched :one

*/
