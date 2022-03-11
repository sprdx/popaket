package routes

import (
	"popaket/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController *users.UserController
}

func (c ControllerList) InitRoutes(e *echo.Echo) {
	e.Debug = true

	// e.POST("/login", controllers.LoginUserController)
	e.POST("/users", c.UserController.Register)

	// r := e.Group("/jwt")
	// r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	// r.GET("/users/:id", controllers.GetUserByIdController)
}
