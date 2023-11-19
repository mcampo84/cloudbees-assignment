package model

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model

	User          *User
	From          string
	To            string
	PurchasePrice float32
}

func NewTicket(user *User, from string, to string, purchasePrice float32) *Ticket {
	return &Ticket{User: user, From: from, To: to, PurchasePrice: purchasePrice}
}
