-- name: CreatePost :many

INSERT INTO posts (id,created_at,updated_at,title,url,description,published_at,feed_id)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
    returning *;


-- name: GetUserPosts :many
SELECT * from posts where feed_id IN
(SELECT feed_id from follows WHERE user_id = $1) LIMIT $2;
