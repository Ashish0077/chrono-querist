package models

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model

	VideoId      string    `gorm:"column:video_id;uniqueIndex;not null;type:text"`
	Title        string    `gorm:"column:title;not null;type:text"`
	Description  string    `gorm:"column:description;type:text"`
	Thumbnail    string    `gorm:"column:thumbnail;type:text"`
	ChannelTitle string    `gorm:"column:channel_title;type:text"`
	PublishedAt  time.Time `gorm:"column:published_at;not null;type:timestamp"`
}

func (Video) TableName() string {
	return "videos"
}

// func NewVideo(title string, description string, url string, duration int, publishedAt time.Time) Video {
// 	return Video{
// 		uuid:        uuid.NewString(),
// 		title:       title,
// 		description: description,
// 		duration:    duration,
// 		url:         url,
// 		publishedAt: publishedAt,
// 		createdAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}
// }
