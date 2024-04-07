package API

import "chrono-querist/internal/core/models"

type VideosPort interface {
	Search(query string) ([]models.Video, error)
	List(page_number int32, offset int32) ([]models.Video, error)
}
