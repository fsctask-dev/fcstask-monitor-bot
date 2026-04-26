package server

import (
	"fcstask-monitor-bot/internal/bot"
	db "fcstask-monitor-bot/internal/db"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func HandleWebhook(bot *bot.Bot) fiber.Handler {
	return adaptor.HTTPHandler(bot.WebhookHandler())
}

func HandleAlert(bot *bot.Bot) fiber.Handler {
	return func(c *fiber.Ctx) error {

		alert, err := ParseAlert(c.Body())
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		users, err := db.GetAllUsers()
		if err != nil {
			log.Printf("[%s][ERROR]: %v", time.Now(), err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		alertText := FormatAlertText(alert)
		SendAlertToUsers(bot, users, alertText)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Alert sent"})
	}
}
