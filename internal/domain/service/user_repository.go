//go:generate mockgen -source=user_repository.go -destination=user_repository.mockgen.go -package=service

package service

type UserRepository interface {
	GetByID(id uint) (User, error)
	FindOrCreate(firstName string, lastName string, email string) User
}
