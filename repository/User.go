package repository

import (
	"github.com/it-02/dormitory/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUser interface {
	AddUser(name string, pass string, isAdmin bool) (db.User, error)
	GetUserByUsername(username string, user *db.User) error
	GetUsersAmount() (int, error)
	UserExists(name string) bool
	IsUserAdmin(uuid string) (bool, error)
}

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) IUser {
	return &User{db: db}
}

func (this User) AddUser(name string, pass string, isAdmin bool) (db.User, error) {
	// add user to db
	user := db.User{
		Username: name,
		Password: pass,
		IsAdmin:  isAdmin,
		UUID:     uuid.New().String(),
	}
	this.db.Create(&user)
	return user, nil
}

func (this User) GetUserByUsername(username string, user *db.User) error {
	err := this.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (this User) GetUsersAmount() (int, error) {
	var users []db.User
	err := this.db.Find(&users).Error
	if err != nil {
		return 0, err
	}
	return len(users), nil
}

func (this User) UserExists(name string) bool {
	var users []db.User
	this.db.Where("username = ?", name).Find(&users)
	return len(users) > 0
}

func (this User) IsUserAdmin(uuid string) (bool, error) {
	var user db.User
	err := this.db.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		return false, err
	}
	return user.IsAdmin, nil
}
