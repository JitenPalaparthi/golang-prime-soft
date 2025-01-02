package handlers

import (
	"errors"
	"fmt"
	"http-echo-demo/interfaces"
	"http-echo-demo/messages"
	"http-echo-demo/models"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	_ "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Userhandler struct {
	IUser       interfaces.IUser
	UserMessage *messages.UserMessage
}

func NewUser(iUser interfaces.IUser, userMessage *messages.UserMessage) *Userhandler {
	return &Userhandler{IUser: iUser, UserMessage: userMessage}
}

func (u *Userhandler) Login(secretKey string) func(c echo.Context) error {
	return func(c echo.Context) error {
		loginUser := new(models.LoginUser)
		err := c.Bind(loginUser)
		if err != nil {
			log.Println(err.Error())
			return errors.New("something went wrong")
		}

		err = loginUser.Validate()
		if err != nil {
			return err
		}

		ok, err := u.IUser.Login(loginUser)
		if err != nil {
			fmt.Println("---in handler", err.Error())
			return err
		}

		if ok {
			claims := jwt.MapClaims{
				"email": loginUser.Email,
				"exp":   time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			str, _ := token.SignedString([]byte(secretKey))
			c.String(200, str)
			return nil
		}
		return nil
	}

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
	u.UserMessage.Value <- user.ToBytes()
	return c.JSON(http.StatusCreated, user)
}

func (u *Userhandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	user, err := u.IUser.GetByID(id)
	if err != nil {
		fmt.Println("------>", err.Error(), user)
		log.Println(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return echo.NewHTTPError(http.StatusNoContent, "record not found")
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "something went wrong")
		}
	}
	return c.JSON(http.StatusOK, user)
}
