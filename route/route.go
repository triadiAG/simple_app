package route

import (
	"myApp/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// Auth
	e.POST("/auth/login", controller.LoginController)

	// Users
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUsersController)
	e.PUT("/users/:id", controller.UpdateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)

	// Books
	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBukuController)
	e.POST("/books", controller.CreateBooksController)
	e.PUT("/books/:id", controller.UpdateBukuController)
	e.DELETE("/books/:id", controller.DeleteBukuController)

	return e
}
