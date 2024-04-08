package models

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model

	VideoId      string    `gorm:"column:video_id;index"`
	Title        string    `gorm:"column:title;index:index_on_title_and_description;class:FULLTEXT;not null"`
	Description  string    `gorm:"column:description;index:index_on_title_and_description;class:FULLTEXT"`
	Thumbnail    string    `gorm:"column:thumbnail"`
	ChannelTitle string    `gorm:"column:channel_title"`
	PublishedAt  time.Time `gorm:"column:published_at;not null;type:timestamp"`
}

func (Video) TableName() string {
	return "videos"
}
