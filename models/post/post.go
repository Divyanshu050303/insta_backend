package post

import (
	"time"

	"gorm.io/gorm"
)

type PostModel struct {
	Id           string    `json:"id" gorm:"type:uuid;primary_key"`
	UserId       string    `json:"userid" gorm:"type:uuid;not null"`
	MediaURL     string    `json:"mediaurl" gorm:"not null"`
	MediaType    string    `json:"mediatype" gorm:"not null"`
	Caption      *string   `json:"caption"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	LikeCount    int       `json:"likecount"`
	CommentCount int       `json:"commentcount"`
}

func MigratePost(db *gorm.DB) error {
	return db.AutoMigrate(&PostModel{})
}
