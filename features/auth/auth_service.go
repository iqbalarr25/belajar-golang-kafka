package auth

import (
	"BelajarKafka/config"
	"BelajarKafka/enums"
	"BelajarKafka/helper"
	"BelajarKafka/helper/auth_helper"
	"BelajarKafka/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	Con *gorm.DB
}

func (service *AuthService) Login(request *LoginRequest, user *models.User) error {
	result := service.Con.Model(user).First(user, "email = ?", request.Email)

	if result.Error != nil {
		err := fmt.Errorf("wrong credentials")
		helper.Exception(err)
		return err
	}

	resultChan := make(chan error)

	go func() {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		resultChan <- err
	}()
	err := <-resultChan

	if result.Error != nil || err != nil {
		err = fmt.Errorf("wrong credentials")
		helper.Exception(err)
		return err
	}

	return nil
}

func (service *AuthService) Register(request *RegisterRequest, user *models.User) error {
	var err error
	if request.Password != request.ConfirmationPassword {
		err = fmt.Errorf("confirmation password do not match")
	}
	if err != nil {
		helper.Exception(err)
		return err
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		helper.Exception(err)
		return err
	}

	*user = models.User{
		Email:    request.Email,
		Password: string(bytes),
		Name:     request.Name,
		Role:     enums.Member,
	}

	result := service.Con.Create(&user)
	if result.Error != nil {
		helper.Exception(result.Error)
		return result.Error
	}

	return nil
}

func (service *AuthService) CreateToken(user *models.User) (string, error) {
	token, err := auth_helper.CreateToken(user)
	if err != nil {
		err = fmt.Errorf("failed to create token")
		return "", err
	}

	return token, nil
}

func (service *AuthService) GetProfile(user *models.User) error {
	result := service.Con.First(&user, "id = ?", user.ID)
	if result.Error != nil {
		helper.Exception(result.Error)
		return result.Error
	}

	return nil
}

func (service *AuthService) UpdateProfile(request *UpdateProfileRequest, user *models.User) error {
	result := service.Con.First(&user, "id = ?", user.ID)
	if result.Error != nil {
		helper.Exception(result.Error)
		return result.Error
	}

	user.Email = request.Email
	user.Name = request.Name

	result = service.Con.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewAuthService() Service {
	return &AuthService{config.CreateDBConnection()}
}
