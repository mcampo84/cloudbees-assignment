//go:generate mockgen -source=user.go -destination=user.mockgen.go -package=service

package service

type User interface {
	GetFirstName() string
	GetLastName() string
	GetEmail() string
}
