-- +goose Up
ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP;

-- adding index for faster lookups/sorting
CREATE INDEX idx_feeds_last_fetched_at ON feeds(last_fetched_at);


-- +goose Down
DROP INDEX IF EXISTS idx_feeds_last_fetched_at;
ALTER TABLE feeds DROP COLUMN last_fetched_at;