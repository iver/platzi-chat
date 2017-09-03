package models

import "log"

type User struct {
	Name     string
	Fullname string
}

func (u *User) User() {
	log.Printf("Desde User.User()")
}
