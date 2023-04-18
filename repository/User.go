package repository

import (
	"github.com/it-02/dormitory/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (u *User) AddUser(name string, pass string, isAdmin bool) (db.User, error) {
	// add user to db
	user := db.User{
		Username: name,
		Password: pass,
		IsAdmin:  isAdmin,
		UUID:     uuid.New().String(),
	}
	u.db.Create(&user)
	return user, nil
}

func (u *User) GetUserByUsername(username string, user *db.User) error {
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetUsersAmount() (int, error) {
	var users []db.User
	err := u.db.Find(&users).Error
	if err != nil {
		return 0, err
	}
	return len(users), nil
}

func (u *User) UserExists(name string) bool {
	var users []db.User
	u.db.Where("username = ?", name).Find(&users)
	return len(users) > 0
}

func (u *User) IsUserAdmin(uuid string) (bool, error) {
	var user db.User
	err := u.db.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		return false, err
	}
	return user.IsAdmin, nil
}
