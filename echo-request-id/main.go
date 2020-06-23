package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/qushot/go-tips/echo-request-id/middlewares"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middlewares.RequestIDInLogger())

	e.GET("/", func(c echo.Context) error {
		c.Logger().Info("Hello!")
		return c.String(http.StatusOK, "Hello!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
