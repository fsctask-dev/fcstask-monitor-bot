package bot

import (
	"context"
	"net/http"

	config "fcstask-monitor-bot/internal/config"
	"fcstask-monitor-bot/internal/grafana"

	gotgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func NewBot(ctx context.Context, cfg *config.Config) (*Bot, error) {
	var gc *grafana.Client
	if cfg.GrafanaURL != "" && cfg.GrafanaToken != "" {
		gc = grafana.NewClient(cfg.GrafanaURL, cfg.GrafanaToken)
	}

	statusHandler := StatusHandler(gc)

	opts := []gotgbot.Option{
		gotgbot.WithMessageTextHandler("/start", gotgbot.MatchTypeExact, Start),
		gotgbot.WithMessageTextHandler("/status", gotgbot.MatchTypeExact, statusHandler),
		gotgbot.WithMessageTextHandler("📊 Статус", gotgbot.MatchTypeExact, statusHandler),
		gotgbot.WithCallbackQueryDataHandler("dashboard:", gotgbot.MatchTypePrefix, DashboardCallback(gc)),
	}

	bot, err := gotgbot.New(cfg.BotToken, opts...)
	if err != nil {
		return nil, err
	}

	if err := RegisterMyCommands(ctx, bot); err != nil {
		return nil, err
	}

	return &Bot{TgBot: bot}, nil
}

func (bot *Bot) StartPolling(ctx context.Context) {
	go func() {
		bot.TgBot.Start(ctx)
	}()
}

func (bot *Bot) StartWebhook(ctx context.Context) {
	go func() {
		bot.TgBot.StartWebhook(ctx)
	}()
}

func (bot *Bot) WebhookHandler() http.Handler {
	return bot.TgBot.WebhookHandler()
}

func RegisterMyCommands(ctx context.Context, bot *gotgbot.Bot) error {
	_, err := bot.SetMyCommands(ctx, &gotgbot.SetMyCommandsParams{
		Commands: []models.BotCommand{
			{Command: "start", Description: "Начать получать уведомления"},
			{Command: "status", Description: "Скриншоты дашбордов Grafana"},
		},
	})
	return err
}
