package user

import (
	"BelajarKafka/enums"
	"BelajarKafka/helper"
	"BelajarKafka/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UpdateRequest struct {
	Email    string         `json:"email" validate:"required,email"`
	Name     string         `json:"name" validate:"required"`
	Password string         `json:"password"`
	Role     enums.UserRole `json:"role"`
}

type CreateRequest struct {
	Email    string         `json:"email" validate:"required,email"`
	Name     string         `json:"name" validate:"required"`
	Password string         `json:"password" validate:"required"`
	Role     enums.UserRole `json:"role" validate:"required"`
}

var Validator = validator.New()

func CreateValidate(c *fiber.Ctx) error {
	var errors []*helper.ValidatorError
	body := CreateRequest{}
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

func UpdateValidate(c *fiber.Ctx) error {
	var errors []*helper.ValidatorError
	body := UpdateRequest{}
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
	Fetch() ([]models.User, error)
	GetByID(id string, user *models.User) error
	Update(id string, request *UpdateRequest, user *models.User) error
	Store(request *CreateRequest, user *models.User) error
	Delete(id string, user *models.User) error
}

type Controller interface {
	Index(ctx *fiber.Ctx) error
	Store(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Destroy(ctx *fiber.Ctx) error
}
