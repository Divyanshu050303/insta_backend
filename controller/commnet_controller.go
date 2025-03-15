package controller

import (
	"divyanshu050303/insta_backend/helper"
	"divyanshu050303/insta_backend/models"
	"divyanshu050303/insta_backend/models/post"
	"divyanshu050303/insta_backend/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CommentController struct {
	Repo *repository.CommentRepository
}

func (ctrl *CommentController) CreateComment(c *fiber.Ctx) error {
	err := helper.CheckUserIsLoggedInOrNot(c)
	if err != nil {
		helper.ApiResponse(c, http.StatusUnauthorized, "token is missing", nil)
		return err
	}
	commentModel := post.CommentModel{
		Id: uuid.New().String(),
	}
	err = c.BodyParser(&commentModel)
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}
	err = ctrl.Repo.DB.Create(&commentModel).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Could Not create the comment", nil)
		return err
	}

	helper.ApiResponse(c, http.StatusOK, "Comment created successfully", nil)

	return nil
}

func (ctrl *CommentController) GetCommentsByPostId(c *fiber.Ctx) error {
	postId := c.Params("postId")
	err := helper.CheckUserIsLoggedInOrNot(c)
	if err != nil {
		helper.ApiResponse(c, http.StatusUnauthorized, "token is missing", nil)
		return err
	}
	var comments []post.CommentModel
	err = ctrl.Repo.DB.Where("post_id=?", postId).Find(&comments).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}
	var commentsInfo []map[string]interface{}
	for _, comment := range comments {
		var user models.UserModels
		err = ctrl.Repo.DB.Where("user_id=?", comment.UserId).Find(&user).Error
		if err != nil {
			return err
		}
		commentMap := map[string]interface{}{
			"id":        comment.Id,
			"comment":   comment.Comment,
			"createdAt": comment.CreatedAt,
			"updatedAt": comment.UpdatedAt,

			"user": map[string]interface{}{
				"userId":    user.UserId,
				"userName":  user.UserName,
				"userEmail": user.UserEmail,
			},
		}
		commentsInfo = append(commentsInfo, commentMap)
	}
	helper.ApiResponse(c, http.StatusOK, "Comments", commentsInfo)
	return nil
}
