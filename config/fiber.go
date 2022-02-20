package config

import (
	"github.com/gofiber/fiber/v2"
	"simple-fiber-crud/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
