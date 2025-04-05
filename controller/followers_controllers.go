package controller

import (
	"divyanshu050303/insta_backend/helper"
	"divyanshu050303/insta_backend/models"
	"divyanshu050303/insta_backend/models/post"
	"divyanshu050303/insta_backend/repository"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FollowersControllers struct {
	Repo *repository.FollowersRepository
}

func (ctrl *FollowersControllers) CreateFollower(c *fiber.Ctx) error {

	var follower models.Followers
	err := helper.CheckUserIsLoggedInOrNot(c)
	if err != nil {
		helper.ApiResponse(c, http.StatusUnauthorized, "token is missing", nil)
		return err
	}
	err = c.BodyParser(&follower)
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}

	var existingFollower models.UserModels
	err = ctrl.Repo.DB.Where("user_id=?", follower.UserId).Find(&existingFollower).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "User Not Found", nil)
		return err
	}
	err = ctrl.Repo.DB.Where("user_id=?", follower.FollowerId).Find(&existingFollower).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "User Not Found", nil)
		return err
	}

	newFollower := models.Followers{
		Id:         uuid.New().String(),
		UserId:     follower.UserId,
		FollowerId: follower.FollowerId,
		CreatedAt:  time.Now(),
	}
	err = ctrl.Repo.DB.Create(&newFollower).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Could Not create the follower", nil)
		return err
	}
	helper.ApiResponse(c, http.StatusCreated, "Follower created successfully", nil)

	return nil
}
func (ctrl *FollowersControllers) UnfollowUser(c *fiber.Ctx) error {
	var unFollow struct {
		UserId     string `json:"userId"`
		FollowerId string `json:"followerId"`
	}
	err := helper.CheckUserIsLoggedInOrNot(c)
	if err != nil {
		helper.ApiResponse(c, http.StatusUnauthorized, "token is missing", nil)
		return err
	}
	err = c.BodyParser(&unFollow)
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}

	err = ctrl.Repo.DB.Where("user_id=? AND follower_id=?", unFollow.UserId, unFollow.FollowerId).Delete(&models.Followers{}).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Could Not unfollow the user", nil)
		return err
	}

	helper.ApiResponse(c, http.StatusOK, "User unfollowed successfully", nil)

	return nil
}

func (ctrl *FollowersControllers) GetFollowers(c *fiber.Ctx) error {
	userId := c.Params("userId")
	err := helper.CheckUserIsLoggedInOrNot(c)
	if err != nil {
		helper.ApiResponse(c, http.StatusUnauthorized, "token is missing", nil)
		return err
	}
	var followers []models.Followers
	err = ctrl.Repo.DB.Where("user_id=?", userId).Find(&followers).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}

	var followersInfo []map[string]interface{}
	for _, follower := range followers {
		var user models.UserModels
		err = ctrl.Repo.DB.Where("user_id=?", follower.FollowerId).Find(&user).Error
		if err != nil {
			return err
		}
		followerMap := map[string]interface{}{
			"userName":  user.UserName,
			"userEmail": user.UserEmail,
		}
		followersInfo = append(followersInfo, followerMap)
	}

	helper.ApiResponse(c, http.StatusOK, "Followers", followersInfo)

	return nil
}

func (ctrl *FollowersControllers) GetFollowing(c *fiber.Ctx) error {
	userId := c.Params("userId")
	err := helper.CheckUserIsLoggedInOrNot(c)
	if err != nil {
		helper.ApiResponse(c, http.StatusUnauthorized, "token is missing", nil)
		return err
	}
	var followers []models.Followers
	err = ctrl.Repo.DB.Where("follower_id=?", userId).Find(&followers).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}

	var followersInfo []map[string]interface{}
	for _, follower := range followers {
		var user models.UserModels
		err = ctrl.Repo.DB.Where("user_id=?", follower.UserId).Find(&user).Error
		if err != nil {
			return err
		}
		followerMap := map[string]interface{}{
			"userName":  user.UserName,
			"userEmail": user.UserEmail,
		}
		followersInfo = append(followersInfo, followerMap)
	}

	helper.ApiResponse(c, http.StatusOK, "Followers", followersInfo)

	return nil
}
func (ctrl *FollowersControllers) GetUserProfile(c *fiber.Ctx) error {
	id := c.Params("userId")
	err := helper.CheckUserIsLoggedInOrNot(c)
	if err != nil {
		helper.ApiResponse(c, http.StatusUnauthorized, "token is missing", nil)
		return err
	}
	if id == "" {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
	}
	var user models.UserModels
	err = ctrl.Repo.DB.Where("user_id=?", id).Find(&user).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}

	var followerCount int64
	err = ctrl.Repo.DB.Model(&models.Followers{}).Where("user_id=?", id).Count(&followerCount).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}
	var followingCount int64
	err = ctrl.Repo.DB.Model(&models.Followers{}).Where("follower_id=?", id).Count(&followingCount).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}
	var posts []post.PostModel
	err = ctrl.Repo.DB.Where("user_id=?", id).Find(&posts).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return err
	}

	var mediaUrls []string
	for _, post := range posts {
		mediaUrls = append(mediaUrls, post.MediaURL)
	}

	userInfo := map[string]interface{}{
		"id":             user.UserId,
		"userName":       user.UserName,
		"userEmail":      user.UserEmail,
		"followerCount":  followerCount,
		"followingCount": followingCount,
		"mediaUrls":      mediaUrls,
	}
	helper.ApiResponse(c, http.StatusOK, "Successfully fetched user profile", userInfo)

	return nil
}

func (ctrl *FollowersControllers) GetSearchResult(c *fiber.Ctx) error {
	searchKey := c.Query("searchKey")
	err := helper.CheckUserIsLoggedInOrNot(c)
	if err != nil {
		helper.ApiResponse(c, http.StatusUnauthorized, "token is missing", nil)
		return err
	}
	if searchKey == "" {
		helper.ApiResponse(c, http.StatusOK, "SuccessFully fetched user", nil)
		return nil
	}
	var users []models.UserModels
	err = ctrl.Repo.DB.Where("LOWER(user_name) LIKE ?", strings.ToLower(searchKey)+"%").Find(&users).Error
	if err != nil {
		helper.ApiResponse(c, http.StatusInternalServerError, "Internal Server error", nil)
		return nil
	}
	var result []map[string]interface{}
	for _, user := range users {
		result = append(result, map[string]interface{}{
			"id":        user.UserId,
			"userName":  user.UserName,
			"userEmail": user.UserEmail,
		})
	}
	helper.ApiResponse(c, http.StatusOK, "SuccessFully fetched user", result)
	return nil
}
