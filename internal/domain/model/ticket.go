package model

import "gorm.io/gorm"

type Section int

const (
	SectionA Section = iota
	SectionB
)

type Ticket struct {
	gorm.Model

	User          *User
	From          string
	To            string
	PurchasePrice float32
	Section       Section
}

func NewTicket(user *User, from string, to string, purchasePrice float32) *Ticket {
	return &Ticket{User: user, From: from, To: to, PurchasePrice: purchasePrice}
}
