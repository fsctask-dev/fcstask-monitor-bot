package server

import (
	"context"

	"fcstask-monitor-bot/internal/bot"
	"fcstask-monitor-bot/internal/logger"
	user "fcstask-monitor-bot/internal/model"

	gotgbot "github.com/go-telegram/bot"
)

func SendAlertToUsers(ctx context.Context, bot *bot.Bot, users []user.User, alertText string) {
	go func() {
		for _, u := range users {
			select {
			case <-ctx.Done():
				return
			default:
			}
			_, err := bot.TgBot.SendMessage(ctx, &gotgbot.SendMessageParams{
				ChatID:    u.ChatID,
				Text:      alertText,
				ParseMode: "HTML",
			})
			if err != nil {
				logger.Log.Error().Err(err).Int64("chat_id", u.ChatID).Msg("failed to send alert")
			}
		}
	}()
}
