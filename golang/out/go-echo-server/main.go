package main

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

    //todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	// DeleteSecureVariable - Delete a secure variable bag
	e.DELETE("/v1/var/:pathToVariable", c.DeleteSecureVariable)

	// GetSecureVariable - Fetch a secure variables bag
	e.GET("/v1/var/:pathToVariable", c.GetSecureVariable)

	// ListVars - List the variable bags
	e.GET("/v1/vars", c.ListVars)

	// UpsertSecureVariable - Store a secure variable bag
	e.PUT("/v1/var/:pathToVariable", c.UpsertSecureVariable)


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}