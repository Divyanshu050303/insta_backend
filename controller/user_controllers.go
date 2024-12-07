package controller

import (
	"divyanshu050303/insta_backend/helper"
	"divyanshu050303/insta_backend/models"
	"divyanshu050303/insta_backend/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserControllers struct {
	Repo *repository.UserRepository
}

func ApiResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(&fiber.Map{
		"status": statusCode, "message": message, "data": data})
}

func (ctrl *UserControllers) CreateUser(c *fiber.Ctx) error {
	userModel := models.UserModels{
		UserId: uuid.New().String(),
	}
	err := c.BodyParser(&userModel)
	if err != nil {
		ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}
	var existingUser models.UserModels
	err = ctrl.Repo.DB.Where("user_email=?", userModel.UserEmail).Find(&existingUser).Error
	if err != nil {
		ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}
	if existingUser.UserId != "" {
		ApiResponse(c, http.StatusConflict, "User Already Exists", nil)
		return err
	}
	err = ctrl.Repo.DB.Create(&userModel).Error
	if err != nil {
		ApiResponse(c, http.StatusBadRequest, "Could Not create the user", nil)
		return err
	}
	accessToken, refreshToken, err := helper.GenerateToken(userModel)
	if err != nil {
		ApiResponse(c, http.StatusBadRequest, "Could Not generate token", nil)
		return err
	}
	myData := map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}

	ApiResponse(c, http.StatusOK, "User Created Successfully", myData)

	return nil
}
func (ctrl *UserControllers) LoginUser(c *fiber.Ctx) error {
	return nil
}
