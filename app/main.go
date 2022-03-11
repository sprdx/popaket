package main

import (
	"log"
	"os"
	"popaket/app/config"
	"popaket/app/routes"
	"popaket/businesses/users"

	_userController "popaket/controllers/users"
	_userDb "popaket/drivers/databases/users"
	"popaket/drivers/mysql"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Debug = true
	config.LoadEnv()

	// connect to db
	configDB := mysql.ConfigDB{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}
	Conn := configDB.InitDB()
	Conn.Migrator().DropTable(&_userDb.Users{})
	Conn.Debug().AutoMigrate(&_userDb.Users{})

	timeoutContextEnv, _ := strconv.Atoi(os.Getenv("TIMEOUT_CONTEXT"))
	timeoutContext := time.Duration(timeoutContextEnv) * time.Second

	userUsecase := users.NewUsecase(_userDb.NewUserRepository(Conn), timeoutContext)

	userController := _userController.NewUserController(*userUsecase)

	// Routes
	routesInit := routes.ControllerList{
		UserController: userController,
	}
	routesInit.InitRoutes(e)
	log.Println(e.Start(":8080"))
}
