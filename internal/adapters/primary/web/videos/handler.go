package videos

import (
	"chrono-querist/internal/ports/primary/API"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Index(videosService API.VideosPort) fiber.Handler {
	return func(fiberContext *fiber.Ctx) error {
		log.Info("Index Action Called")
		// ctx := fiberContext.Context()

		// TODO: Implement index handler
		// - search videos
		// - List videos in paginated manner

		return fiberContext.JSON("")
	}
}
