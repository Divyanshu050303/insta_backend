package repository

import "gorm.io/gorm"

type CommentRepository struct {
	DB *gorm.DB
}
