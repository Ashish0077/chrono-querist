package videos

import (
	"chrono-querist/internal/core/models"
	"fmt"
	"strings"
)

func (videoRp *Adapter) Index(pageNumber int, limit int, title *string, description *string) ([]*models.Video, error) {
	var videos []*models.Video
	query := videoRp.db

	fmt.Println("Checking title")
	if title != nil {
		arrayOfString := strings.Split(*title, " ")
		fmt.Println(arrayOfString)
		for _, word := range arrayOfString {
			query = query.Where("title LIKE ?", "%"+word+"%")
		}
	}

	fmt.Println("Checking description")
	if description != nil {
		arrayOfString := strings.Split(*title, " ")
		fmt.Println("Checking title")
		for _, word := range arrayOfString {
			query = query.Where("description LIKE ?", "%"+word+"%")
		}
	}
	if pageNumber <= 0 {
		pageNumber = 1
	}

	fmt.Println("offset ", limit*(pageNumber-1))

	err := query.Offset(limit * (pageNumber - 1)).Limit(limit).Order("published_at DESC").Find(&videos).Error
	return videos, err
}
