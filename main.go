package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recoverFiber "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"os"
	"simple-fiber-crud/app"
	"simple-fiber-crud/config"
	"simple-fiber-crud/controller"
	"simple-fiber-crud/exception"
	"simple-fiber-crud/repository"
	"simple-fiber-crud/usecase"
	"time"
)

func init() {
	// load .env
	err := godotenv.Load()
	exception.LogError(err)
}

func main() {
	// setup database
	database, err := config.Connect()

	// setup validator
	validate := validator.New()

	// setup repository
	accountRepository := repository.NewAccountRepository(database)

	// setup use case
	accountUseCase := usecase.NewAccountUseCase(&accountRepository, validate)

	// setup controller
	accountController := controller.NewAccountController(&accountUseCase)

	f := fiber.New(config.NewFiberConfig())

	// Next function for skipping logger
	f.Use(logger.New(logger.Config{
		TimeFormat: time.RFC850,
		TimeZone:   "Asia/Jakarta",
		Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/dashboard"
		},
	}))
	f.Use(cors.New())
	f.Use(recoverFiber.New())

	app.NewRouter(accountController, f)

	port := os.Getenv("PORT")

	url := fmt.Sprintf(":%s", port)

	err = f.Listen(url)
	exception.LogError(err)
}
