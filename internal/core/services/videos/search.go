package videos

import (
	"chrono-querist/internal/core/domain/video"
	"context"
	"fmt"
)

func (s *Service) Search(ctx context.Context, query string) ([]video.Video, error) {
	res, err := s.videoPort.Search(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query videos: %w", err)
	}

	return res, nil
}
