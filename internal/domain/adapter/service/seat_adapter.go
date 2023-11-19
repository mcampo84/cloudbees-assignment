package service

import (
	"github.com/mcampo84/cloudbees-assignment/internal/domain/model"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/repository/in_memory"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
)

type SeatAdapter struct {
	seat     *model.Seat
	userRepo *in_memory.UserRepository
}

// AssignPassenger implements service.Seat.
func (a *SeatAdapter) AssignPassenger(passenger service.User) error {
	passengerModel, err := a.userRepo.GetByID(passenger.GetID())
	if err != nil {
		return err
	}

	return a.seat.SetPassenger(passengerModel)
}

// GetNumber implements service.Seat.
func (a *SeatAdapter) GetNumber() int {
	return a.seat.Number
}

// GetPassenger implements service.Seat.
func (a *SeatAdapter) GetPassenger() service.User {
	if a.seat.Passenger == nil {
		return nil
	}

	return NewUserModelAdapter(a.seat.Passenger)
}

// GetSection implements service.Seat.
func (a *SeatAdapter) GetSection() service.Section {
	return NewSectionAdapter(a.seat.Section, a.userRepo)
}

func NewSeatAdapter(seat *model.Seat, userRepo *in_memory.UserRepository) service.Seat {
	return &SeatAdapter{seat: seat, userRepo: userRepo}
}
