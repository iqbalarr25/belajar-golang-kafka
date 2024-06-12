package auth

import (
	"BelajarKafka/helper"
	"BelajarKafka/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required"`
	Name                 string `json:"name" validate:"required"`
	ConfirmationPassword string `json:"confirmation_password" validate:"required"`
}

type UpdateProfileRequest struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required"`
}

var Validator = validator.New()

func LoginValidate(c *fiber.Ctx) error {
	var errors []*helper.ValidatorError
	body := LoginRequest{}
	err := c.BodyParser(&body)
	if err != nil {
		helper.Exception(err)
		return err
	}

	err = Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el helper.ValidatorError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(helper.Response{
			Code:    fiber.StatusBadRequest,
			Success: false,
			Message: "Validation error",
			Data:    errors,
		})
	}
	return c.Next()
}

func RegisterValidate(c *fiber.Ctx) error {
	var errors []*helper.ValidatorError
	body := RegisterRequest{}
	err := c.BodyParser(&body)
	if err != nil {
		helper.Exception(err)
		return err
	}

	err = Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el helper.ValidatorError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(helper.Response{
			Code:    fiber.StatusBadRequest,
			Success: false,
			Message: "Validation error",
			Data:    errors,
		})
	}
	return c.Next()
}

func UpdateProfileValidate(c *fiber.Ctx) error {
	var errors []*helper.ValidatorError
	body := UpdateProfileRequest{}
	err := c.BodyParser(&body)
	if err != nil {
		helper.Exception(err)
		return err
	}

	err = Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el helper.ValidatorError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(helper.Response{
			Code:    fiber.StatusBadRequest,
			Success: false,
			Message: "Validation error",
			Data:    errors,
		})
	}
	return c.Next()
}

type Service interface {
	Register(request *RegisterRequest, user *models.User) error
	Login(request *LoginRequest, user *models.User) error
	CreateToken(user *models.User) (string, error)
	GetProfile(user *models.User) error
	UpdateProfile(request *UpdateProfileRequest, user *models.User) error
}

type Controller interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	GetProfile(ctx *fiber.Ctx) error
	UpdateProfile(ctx *fiber.Ctx) error
}
