package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, *sql.DB) {
	dbUserName := os.Getenv("DB_USER_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbSource := os.Getenv("DB_SOURCE")

	dsn := fmt.Sprintf("%v:%v@tcp(database:%v)/%v", dbUserName, dbPassword, dbPort, dbSource)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("[Connection], Error in opening db")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("[Connection], Error in setting sqldb")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, sqlDB
}
