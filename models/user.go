package models

import (
	"BelajarKafka/enums"
	"BelajarKafka/helper"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary"`
	Email     string         `json:"email" gorm:"type:varchar(255);uniqueIndex"`
	Name      string         `json:"name" gorm:"type:varchar(255)"`
	Password  string         `json:"password"`
	Role      enums.UserRole `json:"role"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type UserLoginResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	DeletedAt time.Time `json:"deleted_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID        uuid.UUID      `json:"id"`
	Email     string         `json:"email"`
	Name      string         `json:"name"`
	Role      enums.UserRole `json:"role"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func ConvertUserToUserResponse(user *User) UserResponse {
	userResponse := UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		DeletedAt: user.DeletedAt,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return userResponse
}

func (user *User) BeforeCreate(con *gorm.DB) (err error) {
	_, err = helper.IsUnique(User{}, "email", user.Email, con)
	if err != nil {
		helper.Exception(err)
		return err
	}
	if user.Role != enums.Admin && user.Role != enums.Member {
		err = fmt.Errorf("role not valid")
		helper.Exception(err)
		return err
	}
	user.ID = uuid.New()
	return nil
}

func (user *User) BeforeUpdate(con *gorm.DB) (err error) {
	userExist := &User{}

	result := con.Model(userExist).First(userExist, "id = ?", user.ID)
	if result.Error != nil {
		helper.Exception(result.Error)
		return result.Error
	}

	if userExist.Email != user.Email {
		_, err = helper.IsUnique(User{}, "email", user.Email, con)
		if err != nil {
			helper.Exception(err)
			return err
		}
	}

	if user.Role != enums.Admin && user.Role != enums.Member {
		err = fmt.Errorf("role not valid")
		helper.Exception(err)
		return err
	}
	return nil
}
