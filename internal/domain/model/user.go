package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string
	LastName  string
	Email     string
	Tickets   []*Ticket
}

func NewUser(firstName string, lastName string, email string) *User {
	return &User{FirstName: firstName, LastName: lastName, Email: email}
}
