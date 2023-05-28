package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserDbTestSuite struct {
	suite.Suite
}

func (suite *UserDbTestSuite) SetupTest() {
}

func (suite *UserDbTestSuite) Test_EmptyOnInit() {
	db := InitDB()
	var users []User
	db.Find(&users)
	assert.Equal(suite.T(), 0, len(users))
}

func (suite *UserDbTestSuite) Test_AddUser() {
	db := InitDB()
	user := User{
		Username: "username",
		Password: "password",
		IsAdmin:  false,
	}
	db.Create(&user)
	var users []User
	db.Find(&users)
	assert.Equal(suite.T(), 1, len(users))

	user = users[0]
	assert.Equal(suite.T(), "username", user.Username)
	assert.Equal(suite.T(), "password", user.Password)
	assert.Equal(suite.T(), false, user.IsAdmin)

	db.Delete(&user)
}

func (suite *UserDbTestSuite) Test_DeleteUser() {
	db := InitDB()
	user := User{
		Username: "username",
		Password: "password",
		IsAdmin:  false,
	}
	db.Create(&user)
	var users []User
	db.Find(&users)
	assert.Equal(suite.T(), 1, len(users))

	user = users[0]
	db.Delete(&user)

	users = []User{}
	db.Find(&users)
	assert.Equal(suite.T(), 0, len(users))
}

func (suite *UserDbTestSuite) Test_UpdateUser() {
	db := InitDB()
	user := User{
		Username: "username",
		Password: "password",
		IsAdmin:  false,
	}
	db.Create(&user)
	var users []User
	db.Find(&users)
	assert.Equal(suite.T(), 1, len(users))

	user = users[0]
	user.Username = "username2"
	user.Password = "password2"
	user.IsAdmin = true
	db.Save(&user)

	users = []User{}
	db.Find(&users)
	assert.Equal(suite.T(), 1, len(users))

	user = users[0]
	assert.Equal(suite.T(), "username2", user.Username)
	assert.Equal(suite.T(), "password2", user.Password)
	assert.Equal(suite.T(), true, user.IsAdmin)

	db.Delete(&user)
}

func TestUserDbTestSuite(t *testing.T) {
	suite.Run(t, new(UserDbTestSuite))
}
