package routes

import (
	"divyanshu050303/insta_backend/controller"
	"divyanshu050303/insta_backend/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetUpUserRoutes(app *fiber.App, db *gorm.DB) {
	userRepository := &repository.UserRepository{DB: db}
	userController := &controller.UserControllers{Repo: userRepository}
	api := app.Group("/api/user")
	api.Post("/createUser", userController.CreateUser)
	api.Post("/login", userController.LoginUser)

}
