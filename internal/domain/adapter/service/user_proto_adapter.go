package service

import (
	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
	"github.com/mcampo84/cloudbees-assignment/internal/proto/pb"
)

type UserProtoAdapter struct {
	user *pb.User
}

// GetEmail implements service.User.
func (a *UserProtoAdapter) GetEmail() string {
	return a.user.GetEmail()
}

// GetFirstName implements service.User.
func (a *UserProtoAdapter) GetFirstName() string {
	return a.user.GetFirstName()
}

// GetLastName implements service.User.
func (a *UserProtoAdapter) GetLastName() string {
	return a.user.GetLastName()
}

func NewUserProtoAdapter(user *pb.User) service.User {
	return &UserProtoAdapter{user: user}
}
