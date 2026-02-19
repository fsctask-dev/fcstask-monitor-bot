package bot

import (
	"context"
	"fcstask-monitor-bot/internal/bot/handlers"
	config "fcstask-monitor-bot/internal/config"
	"net/http"

	gotgbot "github.com/go-telegram/bot"
)

func NewBot(ctx context.Context, cfg *config.Config) (*Bot, error) {
	opts := []gotgbot.Option{
		gotgbot.WithDefaultHandler(handlers.Default),
		gotgbot.WithMessageTextHandler("/start", gotgbot.MatchTypeExact, handlers.Start),
	}

	bot, err := gotgbot.New(cfg.BotToken, opts...)
	if err != nil {
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
