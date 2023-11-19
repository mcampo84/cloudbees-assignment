package service

import (
	"github.com/mcampo84/cloudbees-assignment/internal/domain/model"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/repository/in_memory"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
)

type TrainAdapter struct {
	train    *model.Train
	userRepo *in_memory.UserRepository
}

// GetID implements service.Train.
func (a *TrainAdapter) GetID() uint {
	return a.train.ID
}

// GetSections implements service.Train.
func (a *TrainAdapter) GetSections() []service.Section {
	sections := []service.Section{}

	for _, section := range a.train.Sections {
		sections = append(sections, NewSectionAdapter(section, a.userRepo))
	}

	return sections
}

func NewTrainAdapter(train *model.Train, userRepo *in_memory.UserRepository) service.Train {
	return &TrainAdapter{train: train, userRepo: userRepo}
}
