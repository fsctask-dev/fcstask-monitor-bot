package bot

import (
	"fcstask-monitor-bot/internal/grafana"

	"github.com/go-telegram/bot/models"
)

func mainKeyboard() *models.ReplyKeyboardMarkup {
	return &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{{Text: "📊 Статус"}},
		},
		ResizeKeyboard: true,
	}
}

func dashboardInlineKeyboard() *models.InlineKeyboardMarkup {
	var row []models.InlineKeyboardButton
	var rows [][]models.InlineKeyboardButton
	for i, d := range grafana.Dashboards {
		row = append(row, models.InlineKeyboardButton{
			Text:         d.Name,
			CallbackData: "dashboard:" + d.UID,
		})
		if len(row) == 2 || i == len(grafana.Dashboards)-1 {
			rows = append(rows, row)
			row = nil
		}
	}
	return &models.InlineKeyboardMarkup{InlineKeyboard: rows}
}
