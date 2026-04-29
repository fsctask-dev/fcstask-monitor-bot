package bot

import (
	"context"
	config "fcstask-monitor-bot/internal/config"
	"net/http"

	gotgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func NewBot(ctx context.Context, cfg *config.Config) (*Bot, error) {
	opts := []gotgbot.Option{
		gotgbot.WithDefaultHandler(Default),
		gotgbot.WithMessageTextHandler("/start", gotgbot.MatchTypeExact, Start),
		gotgbot.WithMessageTextHandler("/stop", gotgbot.MatchTypeExact, Stop),
		gotgbot.WithMessageTextHandler("/status", gotgbot.MatchTypeExact, Status),
		gotgbot.WithMessageTextHandler("/help", gotgbot.MatchTypeExact, Help),
	}

	bot, err := gotgbot.New(cfg.BotToken, opts...)
	if err != nil {
		return nil, err
	}

	if _, err := bot.SetMyCommands(ctx, &gotgbot.SetMyCommandsParams{
		Commands: []models.BotCommand{
			{Command: "start", Description: "✅ Подписаться"},
			{Command: "stop", Description: "❌ Отписаться"},
			{Command: "status", Description: "❔ Статус"},
			{Command: "help", Description: "📋 Помощь"},
		},
	}); err != nil {
		return nil, err
	}

	if _, err := bot.SetWebhook(ctx, &gotgbot.SetWebhookParams{
		URL: cfg.PublicURL + "/webhook",
	}); err != nil {
		return nil, err
	}

	return &Bot{
		TgBot: bot,
	}, nil
}

func (bot *Bot) Start(ctx context.Context) {
	go func() {
		bot.TgBot.StartWebhook(ctx)
	}()
}

func (bot *Bot) WebhookHandler() http.Handler {
	return bot.TgBot.WebhookHandler()
}
