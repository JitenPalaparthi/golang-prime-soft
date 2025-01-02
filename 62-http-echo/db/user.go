package db

import (
	"errors"
	"fmt"
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
	fmt.Println("========", user)
	if tx.Error != nil {
		fmt.Println("--------------->>>>", tx.Error)
		return nil, tx.Error
	}
	return user, nil
}

func (u *UserDB) Login(loginUser *models.LoginUser) (bool, error) {
	var count int64
	err := u.DB.Table("users").Where("email = ? and password = ?", loginUser.Email, loginUser.Password).Count(&count).Error
	fmt.Println(err, count)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, err
		}
	}
	if count > 0 {
		return true, nil
	} else {
		return false, errors.New("unauthenticarted user")
	}
}
