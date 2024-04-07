package videos

import (
	"chrono-querist/internal/core/domain/video"
	"chrono-querist/internal/ports"
	"context"
)

type API interface {
	Search(ctx context.Context, query string) ([]video.Video, error)
}

type Service struct {
	videoPort ports.VideoPort
}

func NewService(videoPort ports.VideoPort) *Service {
	return &Service{
		videoPort: videoPort,
	}
}
