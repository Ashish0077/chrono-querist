package videos

import "chrono-querist/internal/core/models"

func (srv *Service) List(page_number int32, offset int32) ([]models.Video, error) {
	// Use Repository to fetch data from database
	// Construct the array of videos
	// return

	videos := []models.Video{}

	return videos, nil
}
