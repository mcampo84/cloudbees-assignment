package in_memory

import (
	domainError "github.com/mcampo84/cloudbees-assignment/internal/domain/error"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/model"
)

type UserRepository struct {
	maxID uint
	users map[uint]*model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{maxID: 0, users: map[uint]*model.User{}}
}

func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, domainError.NotFound("user %d", id)
	}

	return user, nil
}

func (r *UserRepository) Create(firstName string, lastName string, email string) *model.User {
	user := model.NewUser(firstName, lastName, email)
	// auto-increment the ID
	user.ID = r.generateID()

	r.users[user.ID] = user
	r.maxID = user.ID

	return user
}

func (r *UserRepository) generateID() uint {
	return r.maxID + 1
}
