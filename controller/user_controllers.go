package controller

import (
	"divyanshu050303/insta_backend/repository"

	"github.com/gofiber/fiber/v2"
)

type UserControllers struct {
	Repo *repository.UserRepository
}

func (ctrl *UserControllers) CreateUser(c *fiber.Ctx) error {
	return nil
}
func (ctrl *UserControllers) LoginUser(c *fiber.Ctx) error {
	return nil
}
