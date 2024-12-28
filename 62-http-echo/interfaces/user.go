package interfaces

import "http-echo-demo/models"

type IUser interface {
	Create(user *models.User) (*models.User, error)
	GetByID(id string) (*models.User, error)
}
