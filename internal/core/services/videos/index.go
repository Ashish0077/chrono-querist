package videos

import "chrono-querist/internal/core/models"

func (srv *Service) Index(pageNumber int, limit int, title *string, description *string) ([]*models.Video, error) {
	videos, err := srv.videosRepo.Index(pageNumber, limit, title, description)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
