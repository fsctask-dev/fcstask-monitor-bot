package database

import (
	"fmt"
	"os"

	"fcstask-monitor-bot/internal/logger"
	model "fcstask-monitor-bot/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("connect to database: %w", err)
	}

	if err = DB.AutoMigrate(&model.User{}); err != nil {
		return fmt.Errorf("migrate user model: %w", err)
	}

	logger.Log.Info().Msg("database initialized successfully")
	return nil
}

func GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
