-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS Users_Chats
(
    id SERIAL PRIMARY KEY,
    user_id integer not null,
    chat_id integer not null,

    foreign key (user_id) references Users(id),
    foreign key (chat_id) references Chats(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS Users_Chats;
