package videos

import (
	"database/sql"

	"gorm.io/gorm"
)

type Adapter struct {
	db    *gorm.DB
	sqlDB *sql.DB
}

func NewAdapter(db *gorm.DB, sqlDB *sql.DB) *Adapter {
	return &Adapter{
		db:    db,
		sqlDB: sqlDB,
	}
}
