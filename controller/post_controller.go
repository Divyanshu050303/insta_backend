package controller

import (
	"divyanshu050303/insta_backend/helper"
	"divyanshu050303/insta_backend/repository"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type PostController struct {
	Repo *repository.PostRepository
}

func (ctrl *PostController) CreatePost(c *fiber.Ctx) error {
	userId := c.Params("userId")
	urls, err := helper.UploadPost(userId, c)
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}
	fmt.Println(urls)
	return nil
}
func (ctrl *PostController) GetPostsByUserId(c *fiber.Ctx) error {
	return nil
}
