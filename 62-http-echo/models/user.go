package models

import (
	"encoding/json"
	"errors"
)

type User struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Contact      string `json:"contact"`
	Status       string `json:"status"`
	LastModified int64  `json:"last_modified"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("invalid or empty name")
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
