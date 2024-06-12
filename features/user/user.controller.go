package user

import (
	"BelajarKafka/helper"
	"BelajarKafka/models"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service Service
}

func (controller *UserController) Index(ctx *fiber.Ctx) error {
	users, err := controller.service.Fetch()
	if err != nil {
		helper.Exception(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.Response{
			Code:    fiber.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.Response{
		Code:    fiber.StatusOK,
		Success: true,
		Message: "success",
		Data:    users,
	})
}

func (controller *UserController) Store(ctx *fiber.Ctx) error {
	request := CreateRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		helper.Exception(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.Response{
			Code:    fiber.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
	}

	user := models.User{}
	err = controller.service.Store(&request, &user)
	if err != nil {
		helper.Exception(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.Response{
			Code:    fiber.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(helper.Response{
		Code:    fiber.StatusCreated,
		Success: true,
		Message: "success",
		Data:    user,
	})
}

func (controller *UserController) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := models.User{}
	err := controller.service.GetByID(id, &user)
	if err != nil {
		helper.Exception(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.Response{
			Code:    fiber.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.Response{
		Code:    fiber.StatusOK,
		Success: true,
		Message: "success",
		Data:    user,
	})
}

func (controller *UserController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	request := UpdateRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		helper.Exception(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.Response{
			Code:    fiber.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
	}

	user := models.User{}
	err = controller.service.Update(id, &request, &user)
	if err != nil {
		helper.Exception(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.Response{
			Code:    fiber.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.Response{
		Code:    fiber.StatusOK,
		Success: true,
		Message: "success",
		Data:    user,
	})
}

func (controller *UserController) Destroy(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := models.User{}
	err := controller.service.Delete(id, &user)
	if err != nil {
		helper.Exception(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.Response{
			Code:    fiber.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.Response{
		Code:    fiber.StatusOK,
		Success: true,
		Message: "success",
		Data:    user,
	})
}

func NewUserController() Controller {
	return &UserController{
		NewUserService(),
	}
}
