package model

import (
	"gorm.io/gorm"

	domainError "github.com/mcampo84/cloudbees-assignment/internal/domain/error"
)

type Seat struct {
	gorm.Model

	Section   *Section
	Number    int
	Passenger *User
}

func NewSeat(number int, section *Section) *Seat {
	return &Seat{Number: number, Section: section}
}

func (s *Seat) SetPassenger(user *User) error {
	if s.Passenger != nil {
		return domainError.Unavailable("seat %d in %s is alredy assigned", s.Number, s.Section.Label)
	}

	s.Passenger = user

	return nil
}
