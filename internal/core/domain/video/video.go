package video

import (
	"time"

	"github.com/google/uuid"
)

type Video struct {
	uuid        string
	title       string
	description string
	duration    int
	url         string
	publishedAt time.Time
	createdAt   time.Time
	UpdatedAt   time.Time
}

func NewVideo(title string, description string, url string, duration int, publishedAt time.Time) Video {
	return Video{
		uuid:        uuid.NewString(),
		title:       title,
		description: description,
		duration:    duration,
		url:         url,
		publishedAt: publishedAt,
		createdAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
