package in_memory

import (
	"testing"

	"github.com/stretchr/testify/suite"

	domainError "github.com/mcampo84/cloudbees-assignment/internal/domain/error"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/model"
)

type UserRepositorySuite struct {
	suite.Suite
	repo *UserRepository
}

func (suite *UserRepositorySuite) SetupTest() {
	suite.repo = NewUserRepository()
}

func (suite *UserRepositorySuite) TestGetByID_ExistingUser() {
	user := model.NewUser("John", "Doe", "john@example.com")
	suite.repo.users[1] = user

	foundUser, err := suite.repo.GetByID(1)

	suite.Equal(user, foundUser)
	suite.NoError(err)
}

func (suite *UserRepositorySuite) TestGetByID_NonExistingUser() {
	foundUser, err := suite.repo.GetByID(2)

	suite.Nil(foundUser)
	suite.Equal(domainError.NotFound("user %d", 2), err)
}

func (suite *UserRepositorySuite) TestCreateUser() {
	newUser := suite.repo.Create("Jane", "Doe", "jane@example.com")

	suite.NotNil(newUser)
	suite.Equal("Jane", newUser.FirstName)
	suite.Equal("Doe", newUser.LastName)
	suite.Equal("jane@example.com", newUser.Email)
	suite.Equal(uint(1), newUser.ID)

	secondUser := suite.repo.Create("Jane", "Doe", "jane@example.com")
	suite.Equal(uint(2), secondUser.ID)
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
