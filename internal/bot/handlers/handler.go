package handlers

import (
	"context"
	"time"

	database "fcstask-monitor-bot/internal/db"
	model "fcstask-monitor-bot/internal/models"

	gotgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Default(ctx context.Context, bot *gotgbot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
}

func Start(ctx context.Context, bot *gotgbot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID
	user := model.User{
		ChatID:    chatID,
		CreatedAt: time.Now(),
	}

	if err := database.DB.First(&user, "chat_id = ?", chatID).Error; err == nil {
		bot.SendMessage(ctx, &gotgbot.SendMessageParams{
			ChatID: chatID,
			Text:   "⏹️ Вы уже в списке разработчиков, вы можете получать алёрты",
		})
		return
	}

	res := database.DB.Where(model.User{ChatID: chatID}).FirstOrCreate(&user)
	if res.Error != nil {
		bot.SendMessage(ctx, &gotgbot.SendMessageParams{
			ChatID: chatID,
			Text:   "❌ Не удалось добавить вас в список разработчиков, вы не можете получать алёрты",
		})
		return
	}

	bot.SendMessage(ctx, &gotgbot.SendMessageParams{
		ChatID: chatID,
		Text:   "Привет!\n✅ Вы добавлены в список разработчиков, теперь вы можете получать алёрты",
	})
}
