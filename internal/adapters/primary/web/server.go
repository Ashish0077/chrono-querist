package web

import (
	"chrono-querist/internal/ports/primary/API"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type App struct {
	fiber        *fiber.App
	videoService API.VideosPort
	port         string
}

func NewApp(videoService API.VideosPort) *App {
	app := &App{
		fiber:        fiber.New(),
		videoService: videoService,
		port:         os.Getenv("API_SERVER_PORT"),
	}

	app.fiber.Use(logger.New())
	app.RegisterRoutes()

	return app
}

func (app *App) Run() error {
	return app.fiber.Listen(app.port)
}
