package service

import (
	"fmt"

	"github.com/it-02/dormitory/db"
	"github.com/it-02/dormitory/repository"
)

type IUserService interface {
	RegisterUser(name string, pass string) error
	LoginUser(name string, pass string) (string, error)
	IsUserAdmin(uuid string) bool
}

type UserService struct {
	user_repository repository.IUser
}

func NewUserService(user_repository repository.IUser) IUserService {
	return &UserService{user_repository: user_repository}
}

func (this UserService) RegisterUser(name string, pass string) error {
	users_amount, err := this.user_repository.GetUsersAmount()
	if err != nil {
		return err
	}
	if users_amount == 0 {
		// first user is admin
		_, err := this.user_repository.AddUser(name, pass, true)
		if err != nil {
			return err
		}
	} else {
		if this.user_repository.UserExists(name) {
			return fmt.Errorf("user %s already exists", name)
		} else {
			_, err := this.user_repository.AddUser(name, pass, false)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (this UserService) LoginUser(name string, pass string) (string, error) {
	user := db.User{}
	err := this.user_repository.GetUserByUsername(name, &user)
	if err != nil {
		return "", err
	}
	if user.Password != pass {
		return "", fmt.Errorf("wrong password")
	}
	return user.UUID, nil
}

func (this UserService) IsUserAdmin(uuid string) bool {
	is_admin, err := this.user_repository.IsUserAdmin(uuid)
	if err != nil {
		return false
	}
	return is_admin
}
