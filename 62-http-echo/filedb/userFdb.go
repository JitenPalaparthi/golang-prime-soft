package filedb

import (
	"http-echo-demo/models"
	"math/rand/v2"
	"os"
)

type UserFileDB struct {
	FileName string
}

func NewUserFileDB(fn string) *UserFileDB {
	return &UserFileDB{FileName: fn}
}

func (u UserFileDB) Create(user *models.User) (*models.User, error) {
	f, err := os.OpenFile(u.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	user.Id = uint(rand.IntN(1000))
	_, err = f.Write(user.ToBytes())
	if err != nil {
		return nil, err
	}
	return user, nil
}
