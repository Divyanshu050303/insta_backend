package post

import (
	"time"

	"gorm.io/gorm"
)

type CommentModel struct {
	Id        string    `json:"id" gorm:"type:uuid;primary_key"`
	PostId    string    `json:"postid" gorm:"type:uuid;not null"`
	UserId    string    `json:"userid" gorm:"type:uuid;not null"`
	Comment   string    `json:"comment" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func MigrateComment(db *gorm.DB) error {
	return db.AutoMigrate(&CommentModel{})
}
