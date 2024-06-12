package auth_helper

import (
	"BelajarKafka/config"
	"BelajarKafka/enums"
	"BelajarKafka/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

func CreateToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)

	claims := jwt.MapClaims{
		"exp":   expirationTime.Unix(),
		"iat":   time.Now().Unix(),
		"sub":   user.ID.String(),
		"role":  user.Role,
		"name":  user.Name,
		"email": user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.GetEnv("SECRET_KEY", "secret-key")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthMiddleware(c *fiber.Ctx, rolesAuth []enums.UserRole) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"message": fiber.ErrUnauthorized.Error(),
		})
	}

	tokenString := authHeader[7:]
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"message": fiber.ErrUnauthorized.Error(),
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.GetEnv("SECRET_KEY", "secret-key")), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"message": fiber.ErrUnauthorized.Error(),
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"message": fiber.ErrUnauthorized.Error(),
		})
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"message": fiber.ErrUnauthorized.Error(),
		})
	}

	roleString, ok := claims["role"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"message": fiber.ErrUnauthorized.Error(),
		})
	}

	var role enums.UserRole
	switch roleString {
	case string(enums.Admin):
		role = enums.Admin
	default:
		isAuthAvailable := true
		for _, value := range rolesAuth {
			if string(value) == roleString {
				isAuthAvailable = true
				role = enums.Member
				break
			} else {
				isAuthAvailable = false
			}
		}
		if !isAuthAvailable {
			return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
				"message": "User does not have permission",
			})
		}
	}

	user := models.User{
		ID:    uuid.Must(uuid.Parse(userID)),
		Name:  claims["name"].(string),
		Role:  role,
		Email: claims["email"].(string),
	}

	c.Locals("user", user)

	return c.Next()
}

func GetAuthUser(ctx *fiber.Ctx) (*models.User, error) {
	user, ok := ctx.Locals("user").(models.User)
	if !ok {
		return nil, ctx.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"message": fiber.ErrUnauthorized.Error(),
		})
	}
	return &user, nil
}
