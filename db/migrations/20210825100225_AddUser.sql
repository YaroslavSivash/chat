-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    username varchar unique not null,
    created_at timestamptz DEFAULT now()

);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS users;
