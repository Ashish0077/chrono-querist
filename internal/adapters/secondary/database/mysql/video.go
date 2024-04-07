package mysql

import (
	"chrono-querist/internal/core/domain/video"
	"fmt"
	"time"
)

type VideoRepository struct {
	// db *sql.DB
}

func NewVideoRepository() *VideoRepository {
	return &VideoRepository{
		// db: db,
	}
}

func (vr *VideoRepository) Search(query string) ([]video.Video, error) {
	fmt.Printf("Searching for videos with query: %s\n", query)
	videos := []video.Video{
		video.NewVideo("SearchVideo", "SearchVideo", "https://example.com/searchvideo.mp4", 120, time.Now()),
	}

	return videos, nil
}

func (vr *VideoRepository) Upsert(videos []video.Video) error {
	fmt.Println("Upserting videos...")

	return nil
}
