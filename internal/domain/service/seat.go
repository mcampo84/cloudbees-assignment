//go:generate mockgen -source=seat.go -destination=seat.mockgen.go -package=service

package service

type Seat interface {
	AssignPassenger(passenger User) error
	GetPassenger() User
	GetNumber() int
	GetSection() Section
}
