package post

import "gorm.io/gorm"

func MigratePostData(db *gorm.DB) error {
	if err := MigratePost(db); err != nil {
		return err
	}
	return nil
}
