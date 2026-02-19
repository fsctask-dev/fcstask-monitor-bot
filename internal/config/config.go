package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		return nil, errors.New("BOT_TOKEN environment variable is not set")
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		return nil, errors.New("SERVER_PORT environment variable is not set")
	}

	publicURL := os.Getenv("PUBLIC_URL")
	if publicURL == "" {
		return nil, errors.New("PUBLIC_URL environment variable is not set")
	}

	return &Config{
		BotToken:   botToken,
		ServerPort: serverPort,
		PublicURL:  publicURL,
	}, nil
}
