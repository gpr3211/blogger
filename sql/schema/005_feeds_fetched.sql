-- +goose Up

ALTER TABLE feeds
ADD last_fetch TIMESTAMP;

-- +goose Down

ALTER TABLE feeds
DROP COLUMN last_fetch;

