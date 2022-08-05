package main

import (
	"myApp/route"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "myApp/docs"
)

func main() {
	e := route.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Documentation at http://localhost:8080/swagger/index.html")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Start(":8080")
}
