package service

import (
	"github.com/mcampo84/cloudbees-assignment/internal/domain/model"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
)

type UserModelAdapter struct {
	user *model.User
}

// GetEmail implements service.User.
func (a *UserModelAdapter) GetEmail() string {
	return a.user.Email
}

// GetFirstName implements service.User.
func (a *UserModelAdapter) GetFirstName() string {
	return a.user.FirstName
}

// GetLastName implements service.User.
func (a *UserModelAdapter) GetLastName() string {
	return a.user.LastName
}

func NewUserModelAdapter(user *model.User) service.User {
	return &UserModelAdapter{user: user}
}
