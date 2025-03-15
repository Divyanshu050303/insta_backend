package controller

import (
	"divyanshu050303/insta_backend/helper"
	"divyanshu050303/insta_backend/models/post"
	"divyanshu050303/insta_backend/repository"

	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PostController struct {
	Repo *repository.PostRepository
}

func (ctrl *PostController) CreatePost(c *fiber.Ctx) error {
	userId := c.Params("userId")

	postMode := post.PostModel{UserId: userId,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		LikeCount:    0,
		CommentCount: 0,
		Id:           uuid.New().String()}
	err := c.BodyParser(&postMode)
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}
	err = ctrl.Repo.DB.Create(&postMode).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Could Not create the post", nil)
		return err
	}
	helper.ApiResponse(c, http.StatusOK, "Post created successfully", nil)
	return nil
}
func (ctrl *PostController) UploadPost(c *fiber.Ctx) error {
	userId := c.Params("userId")
	urls, err := helper.UploadPost(userId, c)
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}
	helper.ApiResponse(c, http.StatusOK, "Post created successfully", urls)
	return nil

}
func (ctrl *PostController) GetPostsByUserId(c *fiber.Ctx) error {
	return nil
}
