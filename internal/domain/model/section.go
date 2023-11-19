package model

import (
	"gorm.io/gorm"
)

type Section struct {
	gorm.Model

	Train *Train
	Label string
	Seats map[int]*Seat
}

func NewSection(label string, train *Train) *Section {
	section := &Section{Label: label, Seats: make(map[int]*Seat), Train: train}

	for i := 1; i <= 8; i++ {
		section.Seats[i] = NewSeat(i, section)
	}

	return section
}
