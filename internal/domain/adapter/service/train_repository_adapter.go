package service

import (
	"github.com/mcampo84/cloudbees-assignment/internal/domain/repository/in_memory"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
)

type TrainRepositoryAdapter struct {
	trainRepo *in_memory.TrainRepository
	userRepo  *in_memory.UserRepository
}

// Create implements service.TrainRepository.
func (a *TrainRepositoryAdapter) Create(from string, to string) service.Train {
	return NewTrainAdapter(a.trainRepo.Create(from, to), a.userRepo)
}

// Find implements service.TrainRepository.
func (a *TrainRepositoryAdapter) Find(from string, to string) (service.Train, error) {
	train, err := a.trainRepo.Find(from, to)
	if err != nil {
		return nil, err
	}

	return NewTrainAdapter(train, a.userRepo), nil
}

func NewTrainRepositoryAdapter(trainRepo *in_memory.TrainRepository, userRepo *in_memory.UserRepository) service.TrainRepository {
	return &TrainRepositoryAdapter{trainRepo: trainRepo, userRepo: userRepo}
}
