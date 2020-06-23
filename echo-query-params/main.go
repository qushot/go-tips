package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type query struct {
	Slices []string `query:"slices"`
	Embed
}

type Embed struct {
	Embed string `query:"embed"`
}

type another struct {
	Other string `query:"other"`
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		str := ""

		q := new(query)
		if err := c.Bind(q); err != nil {
			return err
		}

		for i, slice := range q.Slices {
			str += fmt.Sprintf("slice[%d]: %s, ", i, slice)
		}
		str += fmt.Sprintf("embed: %s, ", q.Embed.Embed)

		a := new(another)
		if err := c.Bind(a); err != nil {
			return err
		}

		str += fmt.Sprintf("other: %s, ", a.Other)

		return c.String(http.StatusOK, str)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
