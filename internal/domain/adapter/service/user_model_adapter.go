package service

import (
	"github.com/mcampo84/cloudbees-assignment/internal/domain/model"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
)

type UserModelAdapter struct {
	user *model.User
}

func (a *UserModelAdapter) GetID() uint {
	return a.user.ID
}

func (a *UserModelAdapter) GetEmail() string {
	return a.user.Email
}

func (a *UserModelAdapter) GetFirstName() string {
	return a.user.FirstName
}

func (a *UserModelAdapter) GetLastName() string {
	return a.user.LastName
}

func NewUserModelAdapter(user *model.User) service.User {
	return &UserModelAdapter{user: user}
}
