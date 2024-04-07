package main

import (
	"chrono-querist/internal/adapters/primary/web"
	"chrono-querist/internal/adapters/secondary/database/mysql"
	"chrono-querist/internal/core/services/videos"
	"fmt"
)

func main() {
	// fmt.Println("Hello, World!")

	// Initialize the secondary adapters (Port Implementations)
	videosRepo := mysql.NewVideoRepository()
	fmt.Println("Video Repo done")

	// Initialize the core service layer
	// Business Logic is handled here
	searchService := videos.NewService(videosRepo)
	fmt.Println("Search Service done")

	// Init primary adapters (Port Implementations)
	srv := web.NewApp(searchService, web.WithPort(9000))
	fmt.Println("Web Server done")

	srv.Run()
}
