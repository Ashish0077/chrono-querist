package web

import "chrono-querist/internal/adapters/primary/web/videos"

func (app *App) RegisterRoutes() {
	app.fiber.Get("/find", videos.Search(app.videosAPI))
}
