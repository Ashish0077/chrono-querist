package videos

import (
	"chrono-querist/internal/core/models"

	"github.com/gofiber/fiber/v2/log"
)

func (videoRp *Adapter) Search(query string) ([]models.Video, error) {
	// TODO: write logic to search query
	log.Info("Searching videos...")
	return []models.Video{}, nil
}
