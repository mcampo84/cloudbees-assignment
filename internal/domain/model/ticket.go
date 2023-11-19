package model

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model

	User          *User
	Train         *Train
	PurchasePrice float32
	Section       Section
}

func NewTicket(user *User, train *Train, purchasePrice float32) *Ticket {
	return &Ticket{User: user, Train: train, PurchasePrice: purchasePrice}
}
