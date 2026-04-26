package server

import (
	"context"
	"fcstask-monitor-bot/internal/bot"
	user "fcstask-monitor-bot/internal/model"
	"log"
	"time"

	gotgbot "github.com/go-telegram/bot"
)

func SendAlertToUsers(bot *bot.Bot, users []user.User, alertText string) {
	go func() {
		for _, user := range users {
			_, err := bot.TgBot.SendMessage(context.Background(), &gotgbot.SendMessageParams{
				ChatID: user.ChatID,
				Text:   alertText,
			})
			if err != nil {
				log.Printf("[%s][ERROR]: %v, chat_id: %d", time.Now(), err, user.ChatID)
			}
		}
	}()
}
