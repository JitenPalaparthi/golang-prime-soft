package db

import (
	"http-echo-demo/models"

	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{DB: db}
}

func (u *UserDB) Create(user *models.User) (*models.User, error) {
	u.DB.AutoMigrate(&models.User{})
	tx := u.DB.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *UserDB) GetByID(id string) (*models.User, error) {
	user := new(models.User)
	tx := u.DB.First(user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
