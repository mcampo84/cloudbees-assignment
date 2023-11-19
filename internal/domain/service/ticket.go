//go:generate mockgen -source=ticket.go -destination=ticket.mockgen.go -package=service

package service

type Ticket interface {
	GetID() uint
	GetUser() User
	GetFrom() string
	GetTo() string
	GetPurchasePrice() float32
}
