package bot

import (
	"context"
	config "fcstask-monitor-bot/internal/config"
	"net/http"

	gotgbot "github.com/go-telegram/bot"
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
