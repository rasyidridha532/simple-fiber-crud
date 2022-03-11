package exception

import (
	"github.com/gofiber/fiber/v2"
	"simple-fiber-crud/config"
)

func LogError(err error) {
	if err != nil {
		logger := config.Logger()
		logger.Error(err.Error())
		return
	}
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "BAD REQUEST",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":  "INTERNAL SERVER ERROR",
		"message": err.Error(),
	})
}
