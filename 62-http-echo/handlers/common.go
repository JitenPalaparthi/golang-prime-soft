package handlers

import "github.com/labstack/echo/v4"

func Root(c echo.Context) error {
	c.JSON(200, map[string]string{"message": "Hello Primesoft"})
	return nil
}

func Ping(c echo.Context) error {
	c.JSON(200, map[string]string{"reply": "Pong"})
	return nil
}

func Health(c echo.Context) error {
	c.JSON(200, map[string]string{"status": "Ok"})
	return nil
}
