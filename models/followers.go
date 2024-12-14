package models

import (
	"time"

	"gorm.io/gorm"
)

type Followers struct {
	Id         string    `json:"id" gorm:"type:uuid;primary_key"`
	UserId     string    `json:"userid" gorm:"type:uuid;not null"`
	FollowerId string    `json:"followed" gorm:"type:uuid;not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func MigrateFollowers(db *gorm.DB) error {
	return db.AutoMigrate(&Followers{})
}
