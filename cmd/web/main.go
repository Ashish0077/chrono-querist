package main

import (
	"chrono-querist/internal/adapters/primary/web"
	videosRepository "chrono-querist/internal/adapters/secondary/repository/videos"
	videosService "chrono-querist/internal/core/services/videos"
)

func main() {
	// Initialize Repository
	videosRepo := videosRepository.NewAdapter()

	// Initialize Service
	videosService := videosService.NewService(videosRepo)

	// Initialize Primary Driving Adapter
	web.NewApp(videosService).Run()
}
