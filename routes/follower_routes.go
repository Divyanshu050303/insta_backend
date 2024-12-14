package routes

import (
	"divyanshu050303/insta_backend/controller"
	"divyanshu050303/insta_backend/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupFollowerRoutes(app *fiber.App, db *gorm.DB) {
	followersRepository := &repository.FollowersRepository{DB: db}
	followersController := &controller.FollowersControllers{Repo: followersRepository}
	api := app.Group("/api/followers")
	api.Post("/createFollower", followersController.CreateFollower)
	api.Delete("/unFollowUser", followersController.UnfollowUser)
	api.Get("/getFollowers/:userId", followersController.GetFollowers)
	api.Get("/getFollowing/:userId", followersController.GetFollowing)
}
