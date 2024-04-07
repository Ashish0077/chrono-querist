package mysql

import (
	"chrono-querist/internal/core/domain/video"
	"database/sql"
)

type VideoRepository struct {
	db *sql.DB
}

func NewVideoRepository(db *sql.DB) *VideoRepository {
	return &VideoRepository{
		db: db,
	}
}

func (vr *VideoRepository) Search(query string) ([]video.Video, error) {
	// Search Query

	return nil, nil
}

func (vr *VideoRepository) Upsert(videos []video.Video) error {
	// Bulk Upsert Query

	return nil
}
