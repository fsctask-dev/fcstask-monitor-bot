package bot

import (
	"context"
	"log"
	"time"

	database "fcstask-monitor-bot/internal/db"
	model "fcstask-monitor-bot/internal/model"

	gotgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"gorm.io/gorm"
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

	res := database.DB.Where("chat_id = ?", chatID).FirstOrCreate(&user)
	switch {
	case res.Error != nil:
		log.Printf("[%s][ERROR]: %v, chat_id: %d", time.Now(), res.Error, chatID)
	case res.RowsAffected == 0:
		bot.SendMessage(ctx, &gotgbot.SendMessageParams{
			ChatID: chatID,
			Text:   "⏹️ Вы уже подписаны на алёрты",
		})
	default:
		bot.SendMessage(ctx, &gotgbot.SendMessageParams{
			ChatID: chatID,
			Text:   "✅Вы подписались на алёрты",
		})
	}
}

func Stop(ctx context.Context, bot *gotgbot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID

	res := database.DB.Delete(&model.User{}, "chat_id = ?", chatID)
	switch {
	case res.Error != nil:
		log.Printf("[%s][ERROR]: %v, chat_id: %d", time.Now(), res.Error, chatID)
	case res.RowsAffected == 0:
		bot.SendMessage(ctx, &gotgbot.SendMessageParams{
			ChatID: chatID,
			Text:   "⏹️ Вы не были подписаны на алёрты",
		})
	default:
		bot.SendMessage(ctx, &gotgbot.SendMessageParams{
			ChatID: chatID,
			Text:   "❌ Вы отписались от алёртов",
		})
	}
}

func Status(ctx context.Context, bot *gotgbot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID

	res := database.DB.First(&model.User{}, "chat_id = ?", chatID)
	switch {
	case res.Error == gorm.ErrRecordNotFound:
		bot.SendMessage(ctx, &gotgbot.SendMessageParams{
			ChatID: chatID,
			Text:   "❌ Вы не подписаны на алёрты",
		})
	case res.Error != nil:
		log.Printf("[%s][ERROR]: %v, chat_id: %d", time.Now(), res.Error, chatID)
	default:
		bot.SendMessage(ctx, &gotgbot.SendMessageParams{
			ChatID: chatID,
			Text:   "✅ Вы подписаны на алёрты",
		})
	}
}

func Help(ctx context.Context, bot *gotgbot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	bot.SendMessage(ctx, &gotgbot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "/start  — ✅ Подписаться на алёрты\n" +
			"/stop   — ❌ Отписаться от алёртов\n" +
			"/status — ❔ Проверить статус\n" +
			"/help   — 📋 Список команд",
	})
}
