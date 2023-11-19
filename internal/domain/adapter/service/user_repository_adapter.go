package service

import (
	"github.com/mcampo84/cloudbees-assignment/internal/domain/repository/in_memory"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
)

type UserRepositoryAdapter struct {
	userRepo *in_memory.UserRepository
}

// FindOrCreate implements service.UserRepository.
func (a *UserRepositoryAdapter) FindOrCreate(firstName string, lastName string, email string) service.User {
	user := a.userRepo.FindOrCreate(firstName, lastName, email)

	return NewUserModelAdapter(user)
}

// GetByID implements service.UserRepository.
func (a *UserRepositoryAdapter) GetByID(id uint) (service.User, error) {
	user, err := a.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return NewUserModelAdapter(user), nil
}

func NewUserRepositoryAdapter(userRepo *in_memory.UserRepository) service.UserRepository {
	return &UserRepositoryAdapter{userRepo: userRepo}
}
