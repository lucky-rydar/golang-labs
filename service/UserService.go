package service

import (
	"fmt"

	"github.com/it-02/dormitory/db"
	"github.com/it-02/dormitory/repository"
)

func RegisterUser(name string, pass string) error {
	users_amount, err := repository.GetUsersAmount()
	if err != nil {
		return err
	}
	if users_amount == 0 {
		// first user is admin
		_, err := repository.AddUser(name, pass, true)
		if err != nil {
			return err
		}
	} else {
		if repository.UserExists(name) {
			return fmt.Errorf("user %s already exists", name)
		} else {
			_, err := repository.AddUser(name, pass, false)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func LoginUser(name string, pass string) (string, error) {
	user := db.User{}
	err := repository.GetUserByUsername(name, &user)
	if err != nil {
		return "", err
	}
	if user.Password != pass {
		return "", fmt.Errorf("wrong password")
	}
	return user.UUID, nil
}

func IsUserAdmin(uuid string) bool {
	is_admin, err := repository.IsUserAdmin(uuid)
	if err != nil {
		return false
	}
	return is_admin
}
