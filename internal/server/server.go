package server

import (
	"context"
	"fcstask-monitor-bot/internal/config"
	database "fcstask-monitor-bot/internal/db"
	"fmt"
	"log"

	"fcstask-monitor-bot/internal/bot"
	model "fcstask-monitor-bot/internal/models"

	gotgbot "github.com/go-telegram/bot"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func NewServer(ctx context.Context, bot *bot.Bot) *Server {
	app := fiber.New()
	app.Post("/webhook", adaptor.HTTPHandler(bot.WebhookHandler()))
	app.Post("/webhook/alerts", func(c *fiber.Ctx) error {
		var payload AlertPayload
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		var users []model.User
		if err := database.DB.Find(&users).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		
		alertText := "🚨 ALERT\n\n"
		for _, alert := range payload.Alerts {
			summary := alert.Annotations["summary"]
			alertText += fmt.Sprintf("• %s\n", summary)
		}

		go func(users []model.User, alertText string) {
			for _, user := range users {
				if _, err := bot.TgBot.SendMessage(context.Background(), &gotgbot.SendMessageParams{
					ChatID: user.ChatID,
					Text:   alertText,
				}); err != nil {
					log.Printf("failed to send alert to user %d: %v", user.ChatID, err)
				}
			}
		}(users, alertText)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Alert sent"})
	})

	return &Server{
		app: app,
	}
}

func (server *Server) Run(ctx context.Context, cfg *config.Config) error {
	errChan := make(chan error, 1)
	go func() {
		errChan <- server.app.Listen(cfg.ServerPort)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return err
		}
		return nil
	case <-ctx.Done():
		return server.app.Shutdown()
	}
}
