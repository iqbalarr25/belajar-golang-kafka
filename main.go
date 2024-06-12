package main

import (
	"BelajarKafka/config"
	"BelajarKafka/database"
	"BelajarKafka/enums"
	"BelajarKafka/features/auth"
	"BelajarKafka/features/user"
	"BelajarKafka/helper/auth_helper"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
)

var ctx = context.Background()

func main() {
	config.InitEnv()
	config.InitDatabase()
	config.InitCache(ctx)
	database.MigrateDatabase()
	config.InitScheduler()

	app := fiber.New()
	app.Use(
		logger.New(),
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
		}),
	)
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	v1 := app.Group("/api/v1")

	auths := v1.Group("/auth")
	auths.Post("/login", auth.LoginValidate, auth.NewAuthController().Login)
	auths.Post("/register", auth.RegisterValidate, auth.NewAuthController().Register)

	profiles := auths.Group("/profile")
	profiles.Use(func(c *fiber.Ctx) error {
		return auth_helper.AuthMiddleware(c, []enums.UserRole{enums.Admin, enums.Member})
	})
	profiles.Get("/", auth.NewAuthController().GetProfile)
	profiles.Put("/", auth.UpdateProfileValidate, auth.NewAuthController().UpdateProfile)

	users := v1.Group("/users")
	users.Use(func(c *fiber.Ctx) error { return auth_helper.AuthMiddleware(c, []enums.UserRole{enums.Admin}) })
	users.Get("/", user.NewUserController().Index)
	users.Post("/", user.CreateValidate, user.NewUserController().Store)
	users.Get("/:id", user.NewUserController().Show)
	users.Put("/:id", user.UpdateValidate, user.NewUserController().Update)
	users.Delete("/:id", user.NewUserController().Destroy)

	err := app.Listen("0.0.0.0:9000")
	if err != nil {
		panic(err)
	}
}
