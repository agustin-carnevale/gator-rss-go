-- +goose Up
CREATE TABLE feeds (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  name TEXT NOT NULL,
  url TEXT NOT NULL UNIQUE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- Create an index on user_id for faster lookups
CREATE INDEX idx_feeds_user_id ON feeds(user_id);


-- +goose Down
DROP INDEX IF EXISTS idx_feeds_user_id;
DROP TABLE feeds;