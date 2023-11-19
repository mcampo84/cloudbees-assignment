//go:generate mockgen -source=section.go -destination=section.mockgen.go -package=service

package service

type Section interface {
	GetLabel() string
	GetSeats() []Seat
}
