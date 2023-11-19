//go:generate mockgen -source=user.go -destination=user.mockgen.go -package=grpc

package grpc

import "github.com/mcampo84/cloudbees-assignment/internal/proto/pb"

type User interface {
	ToProto() *pb.User
}
