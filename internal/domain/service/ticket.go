//go:generate mockgen -source=ticket.go -destination=ticket.mockgen.go -package=service

package service

type Ticket interface {
	GetUser() User
	GetFrom() string
	GetTo() string
	GetPurchasePrice() float32
}
