-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS Users_Chats
(
    id SERIAL PRIMARY KEY,
    id_user integer not null,
    id_chat integer not null,

    foreign key (id_user) references Users(id),
    foreign key (id_chat) references Chats(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS Users_Chats;
