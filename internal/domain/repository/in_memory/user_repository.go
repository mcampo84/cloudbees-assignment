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

func (r *UserRepository) GetByData(firstName string, lastName string, email string) (*model.User, error) {
	for _, user := range r.users {
		if user.FirstName == firstName && user.LastName == lastName && user.Email == email {
			return user, nil
		}
	}

	return nil, domainError.NotFound("user %s %s", firstName, lastName)
}

func (r *UserRepository) FindOrCreate(firstName string, lastName string, email string) *model.User {
	if user, _ := r.GetByData(firstName, lastName, email); user != nil {
		return user
	}

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
