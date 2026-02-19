package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	
	gotgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []gotgbot.Option{
		gotgbot.WithDefaultHandler(handler),
	}

	godotenv.Load()

	token := os.Getenv("BOT_TOKEN")
	bot, err := gotgbot.New(token, opts...)
	if err != nil {
		log.Fatalf("failed to create bot: %v", err)
	}

	url := os.Getenv("URL")
	if _, err := bot.SetWebhook(ctx, &gotgbot.SetWebhookParams{
		URL: url + "/webhook",
	}); err != nil {
		log.Fatalf("failed to set webhook: %v", err)
	}

	go bot.StartWebhook(ctx)

	server := fiber.New()
	server.Get("/webhook", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "webhook"})
	})
	server.Post("/webhook", adaptor.HTTPHandler(bot.WebhookHandler()))

	go func() {
		if err := server.Listen(":" + os.Getenv("SERVER_PORT")); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	server.Shutdown()
}

func handler(ctx context.Context, bot *gotgbot.Bot, update *models.Update) {
	bot.SendMessage(ctx, &gotgbot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
