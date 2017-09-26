package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello world")
}
