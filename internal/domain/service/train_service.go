//go:generate mockgen -source=train_service.go -destination=train_service.mockgen.go -package=service

package service

import (
	"fmt"

	domainError "github.com/mcampo84/cloudbees-assignment/internal/domain/error"
)

type TrainService interface {
	AutoAssignSeat(train Train, passenger User) error
	NewTrain(from string, to string) Train
	InitTrains()
	GetPassengers(train Train) map[string][]User
}

type trainService struct {
	trainRepo      TrainRepository
	sectionService SectionService
}

func NewTrainService(trainRepo TrainRepository, sectionService SectionService) TrainService {
	return &trainService{trainRepo: trainRepo, sectionService: sectionService}
}

func (s *trainService) NewTrain(from string, to string) Train {
	return s.trainRepo.Create(from, to)
}

func (s *trainService) InitTrains() {
	fmt.Println("seeding trains: New York -> Chicago, Chicago -> Los Angeles, Los Angeles -> New York")
	s.NewTrain("New York", "Chicago")
	s.NewTrain("Chicago", "Los Angeles")
	s.NewTrain("Los Angeles", "New York")
}

func (s *trainService) AutoAssignSeat(train Train, passenger User) error {
	seat, err := s.getNextAvailableSeat(train)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("assigning passenger to seat %d in %s", seat.GetNumber(), seat.GetSection().GetLabel()))
	return seat.AssignPassenger(passenger)
}

func (s *trainService) GetPassengers(train Train) map[string][]User {
	panic("unimplemented")
}

func (s *trainService) getNextAvailableSeat(train Train) (Seat, error) {
	section, err := s.getNextAvailableSection(train)
	if err != nil {
		return nil, err
	}

	seats := section.GetSeats()
	for _, seat := range seats {
		if seat.GetPassenger() == nil {
			return seat, nil
		}
	}

	fmt.Println(fmt.Sprintf("no available seats in %s found", section.GetLabel()))
	return nil, domainError.NotFound("the train is fully booked")
}

func (s *trainService) getNextAvailableSection(train Train) (Section, error) {
	sections := s.getAvailableSections(train)
	var nextSection Section
	mostAvailableSeats := 0

	for _, section := range sections {
		availableSeats := 0
		seats := section.GetSeats()
		for _, seat := range seats {
			if seat.GetPassenger() == nil {
				availableSeats += 1
			}
		}
		if availableSeats > mostAvailableSeats {
			nextSection = section
			mostAvailableSeats = availableSeats
		}
	}

	if nextSection == nil {
		fmt.Println("no available sections found")
		return nil, domainError.NotFound("the train is fully booked")
	}

	return nextSection, nil
}

func (s *trainService) getAvailableSections(train Train) []Section {
	allSections := train.GetSections()
	availableSections := []Section{}

	for _, section := range allSections {
		if s.sectionService.GetOpenSeats(section) > 0 {
			availableSections = append(availableSections, section)
		}
	}

	return availableSections
}
