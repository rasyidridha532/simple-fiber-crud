package controller

import "github.com/gofiber/fiber/v2"

type AccountController interface {
	Create(ctx *fiber.Ctx) error
	List(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
}
