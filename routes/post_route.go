package routes

import (
	"divyanshu050303/insta_backend/controller"
	"divyanshu050303/insta_backend/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetUpPostRoute(app *fiber.App, db *gorm.DB) {
	postRepository := &repository.PostRepository{DB: db}
	postController := &controller.PostController{Repo: postRepository}

	api := app.Group("/api/post")
	api.Post("/createPost/:userId", postController.CreatePost)

	api.Get("/getPosts/:userId", postController.GetPostsByUserId)
	api.Post("/uploadPost/:userId", postController.UploadPost)

}
