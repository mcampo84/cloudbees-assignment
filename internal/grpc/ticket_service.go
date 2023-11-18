//go:generate mockgen -source=ticket_service.go -destination=ticket_service.mockgen.go -package=grpc

package grpc

type TicketService interface {
	PurchaseTicket(userID int64, from string, to string, purchasePrice float32) (Ticket, error)
}
