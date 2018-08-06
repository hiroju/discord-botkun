package http

import (
	"net/http"

	"github.com/hiroju/discord-botkun/debo/usecase"
	"github.com/labstack/echo"
)

func Open(c echo.Context) error {
	debo, err := usecase.GetDebo()
	if err != nil {
		c.String(http.StatusNotFound, "failed create")
	}
	err = debo.Open()
	if err != nil {
		c.String(http.StatusNotFound, "failed open")
	}
	return c.String(http.StatusOK, "open")
}

func Close(c echo.Context) error {
	debo, err := usecase.GetDebo()
	if err != nil {
		c.String(http.StatusNotFound, "no debokun")
	}
	debo.Close()
	return c.String(http.StatusOK, "close")
}
