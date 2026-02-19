package config

import (
	"github.com/joho/godotenv"
	"os"
	"errors"
)

func NewConfig(token, url, serverPort string) *Config {
	return &Config{
		Token:      token,
		URL:        url,
		ServerPort: serverPort,
	}
}

func (c *Config) Validate() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return errors.New("BOT_TOKEN environment variable is not set")
	}
	
	url := os.Getenv("URL")
	if url == "" {
		return errors.New("URL environment variable is not set")
	}
	
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		return errors.New("SERVER_PORT environment variable is not set")
	}
	
	return nil
}
