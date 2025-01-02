package models

import (
	"encoding/json"
	"errors"
)

type User struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Contact      string `json:"contact"`
	Status       string `json:"status"`
	LastModified int64  `json:"last_modified"`
}

type LoginUser struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *LoginUser) Validate() error {
	if u.Password == "" {
		return errors.New("invalid or empty password")
	}

	if u.Email == "" {
		return errors.New("invalid or empty email")
	}
	return nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("invalid or empty name")
	}
	if u.Password == "" {
		return errors.New("invalid or empty password")
	}

	if u.Email == "" {
		return errors.New("invalid or empty email")
	}
	return nil
}

func (u *User) ToBytes() []byte {
	bytes, _ := json.Marshal(u)
	return bytes
}
