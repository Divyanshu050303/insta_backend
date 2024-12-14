package repository

import "gorm.io/gorm"

type FollowersRepository struct {
	DB *gorm.DB
}
