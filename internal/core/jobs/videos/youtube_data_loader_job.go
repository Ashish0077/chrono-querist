package youtube

import (
	videosRepository "chrono-querist/internal/adapters/secondary/repository/videos"
	"chrono-querist/internal/core/lib/external_apis/youtube"
	"chrono-querist/internal/core/models"
	dbUtils "chrono-querist/utils/db"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

func YoutubeDataLoaderJob() {
	log.Info("FETCHING YOUTUBE DATE.....")
	// Connect to Database
	db, sqlDB := dbUtils.Connection()
	defer sqlDB.Close()

	// Initialize Repository
	videosRepo := videosRepository.NewAdapter(db, sqlDB)
	videosMap := make(map[string]bool)
	timeAfter := time.Now().AddDate(0, -1, 0)
	apiKeys := strings.Split(os.Getenv("YOUTUBE_API_KEYS"), ",")
	youtubeClient := youtube.NewYoutubeClient(apiKeys)
	data, err := youtubeClient.SearchByQuery("cricket", 50, timeAfter)
	if err != nil {
		log.Error("Error Fetching data from youtube...")
		return
	}

	var dbVideos []*models.Video
	for _, video := range data {
		var dbVideo models.Video
		dbVideo.VideoId = video.Id.VideoId
		dbVideo.Title = video.Snippet.Title
		dbVideo.Description = video.Snippet.Description
		dbVideo.Thumbnail = video.Snippet.Thumbnails.Default.Url
		dbVideo.ChannelTitle = video.Snippet.ChannelTitle
		timeParsed, _ := time.Parse(time.RFC3339, video.Snippet.PublishedAt)
		dbVideo.PublishedAt = timeParsed
		dbVideo.CreatedAt = time.Now()
		dbVideo.UpdatedAt = time.Now()

		if videosMap[dbVideo.VideoId] {
			continue
		}

		videosMap[dbVideo.VideoId] = true

		dbVideos = append(dbVideos, &dbVideo)
	}

	if len(dbVideos) > 0 {
		err = videosRepo.Upsert(dbVideos)
		if err != nil {
			log.Fatalf("Error Saving videos to the database: %v", err)
		}
	}
}
