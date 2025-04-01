package controller

import (
	"divyanshu050303/insta_backend/helper"
	"divyanshu050303/insta_backend/models"
	"divyanshu050303/insta_backend/models/post"
	"divyanshu050303/insta_backend/repository"
	"net/http"
	"strconv"

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

	// Check if user is logged in
	err := helper.CheckUserIsLoggedInOrNot(c)
	if err != nil {
		helper.ApiResponse(c, http.StatusUnauthorized, "token is missing", nil)
		return err
	}

	// Get pagination parameters (default: page=1, limit=10)
	page, _ := strconv.Atoi(c.Query("page", "1"))    // Default to page 1
	limit, _ := strconv.Atoi(c.Query("limit", "10")) // Default to 10 comments per page

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 { // Prevent excessive limits
		limit = 10
	}

	offset := (page - 1) * limit

	// Fetch total comment count for the post
	var totalComments int64
	ctrl.Repo.DB.Model(&post.CommentModel{}).Where("post_id = ?", postId).Count(&totalComments)

	// Fetch paginated comments
	var comments []post.CommentModel
	err = ctrl.Repo.DB.Where("post_id = ?", postId).Limit(limit).Offset(offset).Find(&comments).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}

	// Fetch user details and map the response
	var commentsInfo []map[string]interface{}
	for _, comment := range comments {
		var user models.UserModels
		err = ctrl.Repo.DB.Where("user_id = ?", comment.UserId).First(&user).Error
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

	// Prepare pagination metadata
	totalPages := (int(totalComments) + limit - 1) / limit // Calculate total pages

	pagination := map[string]interface{}{
		"totalComments": totalComments,
		"totalPages":    totalPages,
		"currentPage":   page,
		"limit":         limit,
	}

	// Send response
	responseData := map[string]interface{}{
		"comments":   commentsInfo,
		"pagination": pagination,
	}

	helper.ApiResponse(c, http.StatusOK, "Comments", responseData)
	return nil
}
