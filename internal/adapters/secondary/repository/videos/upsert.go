package videos

import (
	"chrono-querist/internal/core/models"

	"github.com/gofiber/fiber/v2/log"
)

func (videoRp *Adapter) Upsert(videos []models.Video) (string, error) {
	// TODO: write logic to insert data in bulk
	log.Info("Upserting videos...")
	return "Success", nil
}
