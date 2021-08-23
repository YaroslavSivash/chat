package models

import "time"

type ChatEntity struct {
	Id int
	Name string //уникальное имя чата "unique chat name"
	Users []string //список пользователей в чате (отношение многие-ко-многим)
	CreatedAt time.Time
}
