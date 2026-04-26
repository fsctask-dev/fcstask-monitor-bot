package server

import (
	"context"
	"fcstask-monitor-bot/internal/bot"
	"fcstask-monitor-bot/internal/config"

	"github.com/gofiber/fiber/v2"
)

func NewServer(ctx context.Context, bot *bot.Bot) *Server {
	app := fiber.New()
	app.Post("/webhook", HandleWebhook(bot))
	app.Post("/alert", HandleAlert(bot))

	return &Server{
		app: app,
	}
}

func (server *Server) Run(ctx context.Context, cfg *config.Config) error {
	errChan := make(chan error, 1)

	go func() {
		errChan <- server.app.Listen(":" + cfg.ServerPort)
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return server.app.Shutdown()
	}
}
