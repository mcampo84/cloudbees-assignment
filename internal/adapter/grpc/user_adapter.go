package grpc

import (
	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
	"github.com/mcampo84/cloudbees-assignment/internal/grpc"
	"github.com/mcampo84/cloudbees-assignment/internal/proto/pb"
)

type UserAdapter struct {
	user service.User
}

func NewUserAdapter(user service.User) grpc.User {
	return &UserAdapter{user: user}
}

func (a *UserAdapter) ToProto() *pb.User {
	return &pb.User{
		FirstName: a.user.GetFirstName(),
		LastName:  a.user.GetLastName(),
		Email:     a.user.GetEmail(),
	}
}
