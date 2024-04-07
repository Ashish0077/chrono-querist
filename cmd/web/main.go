package main

import (
	"chrono-querist/internal/adapters/primary/web"
	videosRepository "chrono-querist/internal/adapters/secondary/repository/videos"
	videosService "chrono-querist/internal/core/services/videos"
	dbUtils "chrono-querist/utils/db"
)

func main() {
	// Connect to Database
	db, sqlDB := dbUtils.Connection()
	defer sqlDB.Close()

	// Schema Migrations
	dbUtils.RunSchemaMigrations(db)

	// Initialize Repository
	videosRepo := videosRepository.NewAdapter(db, sqlDB)

	// Initialize Service
	videosService := videosService.NewService(videosRepo)

	// Initialize Primary Driving Adapter
	web.NewApp(videosService).Run()
}
