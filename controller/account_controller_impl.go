package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"simple-fiber-crud/entity"
	"simple-fiber-crud/exception"
	"simple-fiber-crud/usecase"
)

func NewAccountController(accountUseCase *usecase.AccountUseCase) AccountController {
	return &AccountControllerImpl{
		AccountUseCase: *accountUseCase,
	}
}

type AccountControllerImpl struct {
	AccountUseCase usecase.AccountUseCase
}

func (a AccountControllerImpl) Create(ctx *fiber.Ctx) error {
	request := entity.AccountRequest{}
	err := ctx.BodyParser(&request)
	exception.LogError(err)

	request.Id = uuid.New().String()

	response := a.AccountUseCase.Create(request)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"data":   response,
	})
}

func (a AccountControllerImpl) List(ctx *fiber.Ctx) error {
	responses := a.AccountUseCase.List()

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"data":   responses,
	})
}

func (a AccountControllerImpl) GetUser(ctx *fiber.Ctx) error {
	response := a.AccountUseCase.User()

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"data":   response,
	})
}

func (a AccountControllerImpl) DeleteUser(ctx *fiber.Ctx) error {
	response := a.AccountUseCase.DeleteUser

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"data":   response,
	})
}
