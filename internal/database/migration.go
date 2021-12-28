package database

import (
	"github.com/jinzhu/gorm"
	"github.com/veezyjay/go-rest/internal/comment"
)

// MigrateDB - migrates the database and creates the comments table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
