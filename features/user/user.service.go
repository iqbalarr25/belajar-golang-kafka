package user

import (
	"BelajarKafka/config"
	"BelajarKafka/enums"
	"BelajarKafka/helper"
	"BelajarKafka/models"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserService struct {
	Con *gorm.DB
}

func (service *UserService) Fetch() ([]models.User, error) {
	var users []models.User
	service.Con.Find(&users)
	return users, nil
}

func (service *UserService) GetByID(id string, user *models.User) error {
	result := service.Con.First(&user, "id = ?", id)
	if result.Error != nil {
		helper.Exception(result.Error)
		return result.Error
	}
	return nil
}

func (service *UserService) Update(id string, request *UpdateRequest, user *models.User) error {
	var password string
	var role enums.UserRole

	result := service.Con.Model(user).First(user, "id = ?", fmt.Sprintf("%v", id))
	if result.Error != nil {
		helper.Exception(result.Error)
		return result.Error
	}

	if request.Password == "" {
		password = user.Password
	} else {
		bytesValue, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
		if err != nil {
			helper.Exception(err)
			return err
		}

		password = string(bytesValue)
	}

	if request.Role == "" {
		role = user.Role
	} else {
		role = request.Role
	}

	*user = models.User{
		ID:       uuid.Must(uuid.Parse(id)),
		Email:    request.Email,
		Name:     request.Name,
		Password: password,
		Role:     role,
	}

	result = service.Con.Save(user)
	if result.Error != nil {
		helper.Exception(result.Error)
		return result.Error
	}

	return nil
}

func (service *UserService) Store(request *CreateRequest, user *models.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		helper.Exception(err)
		return err
	}

	*user = models.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: string(bytes),
		Role:     request.Role,
	}

	result := service.Con.Create(&user)
	if result.Error != nil {
		helper.Exception(result.Error)
		return result.Error
	}
	return nil
}

func (service *UserService) Delete(id string, user *models.User) error {
	service.Con.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&user)
	return nil
}

func NewUserService() Service {
	return &UserService{config.CreateDBConnection()}
}
