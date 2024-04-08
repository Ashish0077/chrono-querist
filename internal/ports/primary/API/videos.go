package API

import "chrono-querist/internal/core/models"

type VideosPort interface {
	Index(page_number int, limit int, title *string, description *string) ([]*models.Video, error)
}
