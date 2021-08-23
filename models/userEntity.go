package models

import "time"

type UserEntity struct {
	Id        int
	UserName  string    //уникальное имя пользователя "unique username"
	CreatedAt time.Time //время создания пользователя "time created username"
}
