package bot

import (
	"bytes"
	"context"
	"strings"
	"time"

	"fcstask-monitor-bot/internal/db"
	"fcstask-monitor-bot/internal/grafana"
	"fcstask-monitor-bot/internal/logger"
	model "fcstask-monitor-bot/internal/model"

	gotgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Start(ctx context.Context, bot *gotgbot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID
	logger.Log.Info().Int64("chat_id", chatID).Msg("start command received")

	user := model.User{
		ChatID:    chatID,
		CreatedAt: time.Now(),
	}
	if err := database.DB.Where("chat_id = ?", chatID).FirstOrCreate(&user).Error; err != nil {
		logger.Log.Error().Err(err).Int64("chat_id", chatID).Msg("failed to save user")
		return
	}

	bot.SendMessage(ctx, &gotgbot.SendMessageParams{ //nolint:errcheck
		ChatID:      chatID,
		Text:        "✅ Вы будете получать уведомления об алертах системы.",
		ReplyMarkup: mainKeyboard(),
	})
}

func StatusHandler(gc *grafana.Client) gotgbot.HandlerFunc {
	return func(ctx context.Context, bot *gotgbot.Bot, update *models.Update) {
		if update.Message == nil {
			return
		}
		chatID := update.Message.Chat.ID

		if gc == nil {
			bot.SendMessage(ctx, &gotgbot.SendMessageParams{ 
				ChatID: chatID,
				Text:   "⚠️ Grafana не настроена (GRAFANA_URL / GRAFANA_TOKEN не заданы).",
			})
			return
		}

		bot.SendMessage(ctx, &gotgbot.SendMessageParams{ 
			ChatID:      chatID,
			Text:        "Выберите дашборд:",
			ReplyMarkup: dashboardInlineKeyboard(),
		})
	}
}

func DashboardCallback(gc *grafana.Client) gotgbot.HandlerFunc {
	return func(ctx context.Context, b *gotgbot.Bot, update *models.Update) {
		query := update.CallbackQuery
		if query == nil {
			return
		}

		b.AnswerCallbackQuery(ctx, &gotgbot.AnswerCallbackQueryParams{
			CallbackQueryID: query.ID,
			Text:            "Загрузка...",
		})

		uid := strings.TrimPrefix(query.Data, "dashboard:")
		chatID := query.From.ID

		imgData, err := gc.RenderDashboard(ctx, uid)
		if err != nil {
			logger.Log.Error().Err(err).Str("uid", uid).Msg("failed to render dashboard")
			b.SendMessage(ctx, &gotgbot.SendMessageParams{
				ChatID: chatID,
				Text:   "❌ Не удалось получить скриншот дашборда.",
			})
			return
		}

		b.SendPhoto(ctx, &gotgbot.SendPhotoParams{
			ChatID: chatID,
			Photo:  &models.InputFileUpload{Filename: "screenshot.png", Data: bytes.NewReader(imgData)},
		})
	}
}
