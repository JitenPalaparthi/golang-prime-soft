package handlers

import (
	"errors"
	"http-echo-demo/interfaces"
	"http-echo-demo/models"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Userhandler struct {
	IUser interfaces.IUser
}

func NewUser(iUser interfaces.IUser) *Userhandler {
	return &Userhandler{IUser: iUser}
}

func (u *Userhandler) Create(c echo.Context) error {
	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		log.Println(err.Error())
		return errors.New("something went wrong")
	}

	err = user.Validate()
	if err != nil {
		return err
	}
	user.Status = "active"
	user.LastModified = time.Now().Unix()

	user, err = u.IUser.Create(user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}
