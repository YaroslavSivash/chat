package models

import "time"

type MessageEntity struct {
	Id        int    //уникальный индентификатор сообщения
	Chat      string //ссылка на индентификатор чата, в который было отправлено сообщение
	Author    string //ссылка на индентификатор отправителя сообщения, отношение многие-к-одному
	Text      string //текст отправленого сообщения
	CreatedAt time.Time
}
