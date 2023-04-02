package repository

import (
	"github.com/it-02/dormitory/db"
	"github.com/google/uuid"
)

func AddUser(name string, pass string, isAdmin bool) (db.User, error) {
	// add user to db
	user := db.User{
		Username: name,
		Password: pass,
		IsAdmin:  isAdmin,
		UUID:     uuid.New().String(),
	}
	db.DB.Create(&user)
	return user, nil
}

func GetUserByUsername(username string, user *db.User) error {
	err := db.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUsersAmount() (int, error) {
	var users []db.User
	err := db.DB.Find(&users).Error
	if err != nil {
		return 0, err
	}
	return len(users), nil
}

func UserExists(name string) bool {
	var users []db.User
	db.DB.Where("username = ?", name).Find(&users)
	return len(users) > 0
}

func IsUserAdmin(uuid string) (bool, error) {
	var user db.User
	err := db.DB.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		return false, err
	}
	return user.IsAdmin, nil
}
