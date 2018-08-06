package main

import (
	"net/http"
	"os"

	debo_http "github.com/hiroju/discord-botkun/debo/delivery/http"
	"github.com/labstack/echo"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/debo/open", debo_http.Open)
	e.GET("/debo/close", debo_http.Close)
	e.Logger.Fatal(e.Start(":" + port))
}
