package repository

import "chrono-querist/internal/core/models"

type VideosPort interface {
	Upsert(videos []*models.Video) error
	Index(page_number int, limit int, title *string, description *string) ([]*models.Video, error)
}
