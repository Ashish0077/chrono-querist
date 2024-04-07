package db

import (
	"chrono-querist/internal/core/models"

	"gorm.io/gorm"
)

func RunSchemaMigrations(db *gorm.DB) {
	var videos models.Video
	db.AutoMigrate(&videos)
}
