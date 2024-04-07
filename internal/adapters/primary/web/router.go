package web

import (
	"chrono-querist/internal/adapters/primary/web/videos"

	"github.com/gofiber/fiber/v2"
)

func registerVideosRoutes(router fiber.Router, app *App) {
	router.Get("/search", videos.Index(app.videoService))
}

func (app *App) RegisterRoutes() {
	registerVideosRoutes(app.fiber.Group("/videos"), app)
}
