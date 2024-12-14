package controller

import (
	"divyanshu050303/insta_backend/helper"
	"divyanshu050303/insta_backend/models"
	"divyanshu050303/insta_backend/repository"

	"strings"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserControllers struct {
	Repo *repository.UserRepository
}

func (ctrl *UserControllers) CreateUser(c *fiber.Ctx) error {
	userModel := models.UserModels{
		UserId: uuid.New().String(),
	}
	err := c.BodyParser(&userModel)
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}
	var existingUser models.UserModels
	err = ctrl.Repo.DB.Where("user_email=?", userModel.UserEmail).Find(&existingUser).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}
	if existingUser.UserId != "" {
		helper.ApiResponse(c, http.StatusConflict, "User Already Exists", nil)
		return err
	}
	err = ctrl.Repo.DB.Create(&userModel).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Could Not create the user", nil)
		return err
	}
	accessToken, refreshToken, err := helper.GenerateToken(userModel)
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Could Not generate token", nil)
		return err
	}
	myData := map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}

	helper.ApiResponse(c, http.StatusOK, "User Created Successfully", myData)

	return nil
}
func (ctrl *UserControllers) LoginUser(c *fiber.Ctx) error {
	userModel := models.UserModels{}
	err := c.BodyParser(&userModel)
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}
	var existingUser models.UserModels
	err = ctrl.Repo.DB.Where("user_email=?", userModel.UserEmail).Find(&existingUser).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}
	if existingUser.UserId == "" {
		helper.ApiResponse(c, http.StatusNotFound, "User Not Found", nil)
		return nil
	}

	if !strings.EqualFold(*existingUser.UserPassword, *userModel.UserPassword) {
		helper.ApiResponse(c, http.StatusUnauthorized, "Invalid Password", nil)
		return nil
	}
	accessToken, refreshToken, err := helper.GenerateToken(existingUser)
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Could Not generate token", nil)
		return err
	}
	myData := map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}
	helper.ApiResponse(c, http.StatusOK, "User Logged In Successfully", myData)
	return nil
}
