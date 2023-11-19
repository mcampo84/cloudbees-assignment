package service

import (
	"github.com/mcampo84/cloudbees-assignment/internal/domain/model"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/repository/in_memory"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
)

type SectionAdapter struct {
	section  *model.Section
	userRepo *in_memory.UserRepository
}

// GetLabel implements service.Section.
func (a *SectionAdapter) GetLabel() string {
	return a.section.Label
}

// GetSeats implements service.Section.
func (a *SectionAdapter) GetSeats() []service.Seat {
	seats := []service.Seat{}
	for _, seat := range a.section.Seats {
		seats = append(seats, NewSeatAdapter(seat, a.userRepo))
	}

	return seats
}

func NewSectionAdapter(section *model.Section, userRepo *in_memory.UserRepository) service.Section {
	return &SectionAdapter{section: section, userRepo: userRepo}
}
