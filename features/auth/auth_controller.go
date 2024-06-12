package auth

import (
	"BelajarKafka/helper"
	"BelajarKafka/helper/auth_helper"
	"BelajarKafka/models"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service Service
}

func (controller *AuthController) Login(ctx *fiber.Ctx) error {
	request := LoginRequest{}
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
	err = controller.service.Login(&request, &user)
	if err != nil {
		helper.Exception(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.Response{
			Code:    fiber.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
	}

	token, err := controller.service.CreateToken(&user)
	if err != nil {
		helper.Exception(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(helper.Response{
			Code:    fiber.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
	}

	userLoginResponse := models.UserLoginResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		Token: token,
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.Response{
		Code:    fiber.StatusOK,
		Success: true,
		Message: "success",
		Data:    userLoginResponse,
	})
}

func (controller *AuthController) Register(ctx *fiber.Ctx) error {
	request := RegisterRequest{}
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
	err = controller.service.Register(&request, &user)
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

func (controller *AuthController) GetProfile(ctx *fiber.Ctx) error {
	user, err := auth_helper.GetAuthUser(ctx)
	if err != nil {
		helper.Exception(err)
		return err
	}

	err = controller.service.GetProfile(user)
	if err != nil {
		helper.Exception(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.Response{
			Code:    fiber.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
	}

	userResponse := models.ConvertUserToUserResponse(user)

	return ctx.Status(fiber.StatusOK).JSON(helper.Response{
		Code:    fiber.StatusOK,
		Success: true,
		Message: "success",
		Data:    userResponse,
	})
}

func (controller *AuthController) UpdateProfile(ctx *fiber.Ctx) error {
	request := UpdateProfileRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		helper.Exception(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.Response{
			Code:    fiber.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
	}

	user, err := auth_helper.GetAuthUser(ctx)
	if err != nil {
		helper.Exception(err)
		return err
	}

	err = controller.service.UpdateProfile(&request, user)

	return ctx.Status(fiber.StatusOK).JSON(helper.Response{
		Code:    fiber.StatusOK,
		Success: true,
		Message: "success",
		Data:    user,
	})
}

func NewAuthController() Controller {
	return &AuthController{
		NewAuthService(),
	}
}
