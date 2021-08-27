-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS message
(
    id SERIAL PRIMARY KEY,
    chat_id integer not null ,
    author_id integer not null ,
    text varchar not null ,
    created_at timestamptz DEFAULT now(),

    foreign key (chat_id) references chats(id),
    foreign key (author_id) references users(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS message;
