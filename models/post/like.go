package post

import (
	"time"

	"gorm.io/gorm"
)

type LikeCountModel struct {
	Id        string    `json:"id" gorm:"type:uuid;primary_key"`
	PostId    string    `json:"postid" gorm:"type:uuid;not null"`
	UserId    string    `json:"userid" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func MigrateLikeCount(db *gorm.DB) error {
	return db.AutoMigrate(&LikeCountModel{})
}
