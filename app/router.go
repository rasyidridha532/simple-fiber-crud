package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"simple-fiber-crud/controller"
	"time"
)

func NewRouter(controller controller.AccountController) (app *fiber.App) {
	app.Get("/dashboard", monitor.New())

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/", func(ctx *fiber.Ctx) error {
		now := time.Now()

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"msg":           "Hello!",
			"response_time": time.Since(now).String(),
		})
	})
	v1.Get("/accounts", controller.List)
	v1.Get("/account", controller.GetUser)
	v1.Post("/account", controller.Create)
	v1.Delete("/account", controller.DeleteUser)

	return app
}
