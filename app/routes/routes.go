package routes

import (
	"os"
	"popaket/app/config"
	"popaket/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	UserController *users.UserController
}

func (c ControllerList) InitRoutes(e *echo.Echo) {
	config.LoadEnv()
	e.Debug = true

	e.POST("/users", c.UserController.Register)
	e.POST("/login", c.UserController.Login)

	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(os.Getenv("KEY_JWT"))))
	r.GET("/users", c.UserController.GetById)
}
