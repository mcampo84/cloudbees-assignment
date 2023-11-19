//go:generate mockgen -source=section_service.go -destination=section_service.mockgen.go -package=service

package service

import (
	"fmt"

	domainError "github.com/mcampo84/cloudbees-assignment/internal/domain/error"
)

type SectionService interface {
	GetOpenSeats(section Section) int
	GetNextOpenSeat(section Section) (Seat, error)
}

type sectionService struct{}

func NewSectionService() SectionService {
	return &sectionService{}
}

func (a *sectionService) GetOpenSeats(section Section) int {
	openSeats := 0
	for _, seat := range section.GetSeats() {
		if seat.GetPassenger() == nil {
			openSeats += 1
		}
	}

	fmt.Println(fmt.Sprintf("%s has %d available seats", section.GetLabel(), openSeats))

	return openSeats
}

func (a *sectionService) GetNextOpenSeat(section Section) (Seat, error) {
	for _, seat := range section.GetSeats() {
		if seat.GetPassenger() == nil {
			return seat, nil
		}
	}

	return nil, domainError.NotFound("no more seats available in section %s", section.GetLabel())
}
