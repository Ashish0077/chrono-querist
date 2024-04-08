package videos

import (
	"chrono-querist/internal/core/models"

	"gorm.io/gorm/clause"
)

func (videoRp *Adapter) Upsert(videos []*models.Video) error {
	return videoRp.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "VideoId"}},
		DoNothing: true,
	}).Create(&videos).Error
}
