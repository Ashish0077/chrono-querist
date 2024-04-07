package videos

import (
	"chrono-querist/internal/core/models"

	"github.com/gofiber/fiber/v2/log"
)

func (videoRp *Adapter) List(page_number int32, offset int32) ([]models.Video, error) {
	// TODO: write logic to list in paginated manner
	log.Info("Listing videos...")
	return []models.Video{}, nil
}
