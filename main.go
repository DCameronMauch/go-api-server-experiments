package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/DCameronMauch/go-api-server-experiments/handlers"
)

func main() {
	server := configure_server()
	configure_routes(server)
	server.Logger.Fatal(server.Start(":8080"))
}

func configure_server() *echo.Echo {
	server := echo.New()

	server.Use(middleware.CSRF())
	server.Use(middleware.Gzip())
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.RemoveTrailingSlash())
	server.Use(middleware.RequestID())
	server.Use(middleware.Secure())

	return server
}

func configure_routes(server *echo.Echo) {
	server.GET("/", handlers.HelloWorld)
}
