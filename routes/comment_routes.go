package routes

import (
	"divyanshu050303/insta_backend/controller"
	"divyanshu050303/insta_backend/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SrtUpCommnetRoutes(app *fiber.App, db *gorm.DB) {
	commentRepository := &repository.CommentRepository{DB: db}
	commentController := &controller.CommentController{Repo: commentRepository}

	api := app.Group("/api/comment")
	api.Post("/createComment", commentController.CreateComment)
	api.Get("/getComments/:postId", commentController.GetCommentsByPostId)
}
