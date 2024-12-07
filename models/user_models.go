package models

import "gorm.io/gorm"

type UserModels struct {
	UserId       string  `json:"userid" gorm:"type:uuid;primary_key"`
	UserEmail    *string `json:"useremail" gorm:"type:varchar(255);not null;unique"`
	UserPassword *string `json:"userpassword" gorm:"type:varchar(255);not null"`
	UserName     *string `json:"username" gorm:"type:varchar(255);not null;unique"`
}

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&UserModels{})
}
