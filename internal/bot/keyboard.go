package bot

import "github.com/go-telegram/bot/models"

func BuildReplyKeyboard() models.ReplyKeyboardMarkup {
	subBtn := models.KeyboardButton{
		Text: "✅ Подписаться",
	}
	unsubBtn := models.KeyboardButton{
		Text: "❌ Отписаться",
	}
	statusBtn := models.KeyboardButton{
		Text: "❔ Статус",
	}
	helpBtn := models.KeyboardButton{
		Text: "📋 Помощь",
	}
	keyboard := [][]models.KeyboardButton{
		{subBtn, unsubBtn},
		{statusBtn, helpBtn},
	}

	MainReplyKeyboard := models.ReplyKeyboardMarkup{
		Keyboard:        keyboard,
		IsPersistent:    true,
		ResizeKeyboard:  true,
		OneTimeKeyboard: false,
		Selective:       false,
	}

	return MainReplyKeyboard
}
