package models

import "log"

type User struct {
	Id       int64
	Username string
	Email    string
}

func (u *User) User() {
	log.Printf("Desde User.User()")
}
