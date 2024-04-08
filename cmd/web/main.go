package main

import (
	"chrono-querist/internal/adapters/primary/web"
	videosRepository "chrono-querist/internal/adapters/secondary/repository/videos"
	videoJobs "chrono-querist/internal/core/jobs/videos"
	videosService "chrono-querist/internal/core/services/videos"
	dbUtils "chrono-querist/utils/db"

	"github.com/robfig/cron/v3"
)

func main() {
	// Connect to Database
	db, sqlDB := dbUtils.Connection()
	defer sqlDB.Close()

	// Schema Migrations
	dbUtils.RunSchemaMigrations(db)

	// Initialize Repository
	videosRepo := videosRepository.NewAdapter(db, sqlDB)

	// Initialize Cron Jobs
	cronJobs := cron.New()
	cronJobs.AddFunc("*/5 * * * *", videoJobs.YoutubeDataLoaderJob)
	cronJobs.Start()

	// Initialize Service
	videosService := videosService.NewService(videosRepo)

	// Initialize Primary Driving Adapter
	web.NewApp(videosService).Run()
}
