package server

import (
	"context"

	"fcstask-monitor-bot/internal/bot"
	db "fcstask-monitor-bot/internal/db"
	"fcstask-monitor-bot/internal/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func HandleWebhook(bot *bot.Bot) fiber.Handler {
	return adaptor.HTTPHandler(bot.WebhookHandler())
}

func HandleAlert(ctx context.Context, bot *bot.Bot) fiber.Handler {
	return func(c *fiber.Ctx) error {
		alert, err := ParseAlert(c.Body())
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		users, err := db.GetAllUsers()
		if err != nil {
			logger.Log.Error().Err(err).Msg("failed to get users from database")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		alertText := FormatAlertText(alert)
		SendAlertToUsers(ctx, bot, users, alertText)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Alert sent"})
	}
}
