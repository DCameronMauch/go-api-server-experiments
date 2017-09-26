package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/DCameronMauch/go-api-server-experiments/handlers"
)

func main() {
	e := configure_server()
	configure_routes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func configure_server() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CSRF())
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())

	return e
}

func configure_routes(e *echo.Echo) {
	e.GET("/", handlers.HelloWorld)
}
