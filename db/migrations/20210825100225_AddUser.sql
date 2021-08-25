-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS Users
(
    id SERIAL PRIMARY KEY,
    username varchar unique not null,
    created_at timestamp

);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS Users;
