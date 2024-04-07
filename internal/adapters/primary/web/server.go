package web

import (
	"chrono-querist/internal/core/services/videos"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	fiber     *fiber.App
	videosAPI videos.API
	port      int
}

func NewApp(videosAPI videos.API, opts ...AppOption) *App {
	s := &App{
		fiber:     fiber.New(),
		videosAPI: videosAPI,
		port:      9000,
	}
	for _, applyOption := range opts {
		applyOption(s)
	}

	s.RegisterRoutes()

	return s
}

func (app *App) Run() error {
	return app.fiber.Listen(fmt.Sprintf(":%d", app.port))
}
