package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func GormOpen() (gormDB *gorm.DB, err error) {
	// dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s connect_timeout=%s sslmode=disable",
	// 	"postgres",
	// 	"Admin123",
	// 	"todo",
	// 	"5432",
	// 	"localhost",
	// 	"300",
	// 	//os.Getenv("PG_TIMEZONE"),
	// )
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s connect_timeout=%s sslmode=disable",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DATABASE"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_CONNECT_TIMEOUT"),
		//os.Getenv("PG_TIMEZONE"),
	)

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db, nil

}
