package web

import (
	"chrono-querist/internal/ports/primary/API"
	"fmt"
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

	app.fiber.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "============>> ${method}#${path} ## QueryParams:${queryParams} ## bodyParams:${bodyParams} ## headers:${reqHeaders}\n",
	}))

	app.RegisterRoutes()

	return app
}

func (app *App) Run() error {
	return app.fiber.Listen(fmt.Sprintf(":%v", app.port))
}
