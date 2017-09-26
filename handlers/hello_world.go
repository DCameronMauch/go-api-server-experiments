package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func HelloWorld(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "hello world")
}
