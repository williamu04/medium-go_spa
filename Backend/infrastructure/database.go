package infrastructure

import (
	"fmt"

	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(dsn string, logger *pkg.Logger) (*gorm.DB, error) {
	logger.Infof("Initializing database connection to PostgreSQL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errorf("Failed to open database connection: %v", err)
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		logger.Errorf("Failed to ping database: %v", err)
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)

	logger.Infof("Database connection established successfully")

	return db, nil
}
