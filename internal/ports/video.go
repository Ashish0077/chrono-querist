package ports

import (
	"chrono-querist/internal/core/domain/video"
)

type VideoPort interface {
	Search(query string) ([]video.Video, error)
	Upsert(videos []video.Video) error
}
