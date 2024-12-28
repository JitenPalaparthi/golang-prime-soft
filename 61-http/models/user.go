package models

import "errors"

type User struct {
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
