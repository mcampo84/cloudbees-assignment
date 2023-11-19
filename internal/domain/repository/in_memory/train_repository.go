package in_memory

import (
	domainError "github.com/mcampo84/cloudbees-assignment/internal/domain/error"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/model"
)

type TrainRepository struct {
	maxID  uint
	trains map[uint]*model.Train
}

func NewTrainRepository() *TrainRepository {
	trains := make(map[uint]*model.Train)

	return &TrainRepository{maxID: 0, trains: trains}
}

func (r *TrainRepository) GetByID(id uint) (*model.Train, error) {
	train, ok := r.trains[id]
	if !ok {
		return nil, domainError.NotFound("train with id %d", id)
	}

	return train, nil
}

func (r *TrainRepository) Find(from string, to string) (*model.Train, error) {
	for _, train := range r.trains {
		if train.From == from && train.To == to {
			return train, nil
		}
	}

	return nil, domainError.NotFound("train from %s to %s not found", from, to)
}

func (r *TrainRepository) Create(from string, to string) *model.Train {
	if existingTrain, _ := r.Find(from, to); existingTrain != nil {
		return existingTrain
	}

	train := model.NewTrain(from, to)
	train.ID = r.generateID()

	r.trains[train.ID] = train
	r.maxID = train.ID

	return train
}

func (r *TrainRepository) generateID() uint {
	return r.maxID + 1
}
