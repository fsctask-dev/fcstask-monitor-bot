package bot

import (
	"context"
	"time"

	"fcstask-monitor-bot/internal/db"
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

	_, err := bot.SendMessage(ctx, &gotgbot.SendMessageParams{
		ChatID: chatID,
		Text:   "✅ Вы будете получать уведомления об алертах системы.",
	})
	logger.Log.Debug().Err(err).Int64("chat_id", chatID).Msg("sent welcome message")
}

