package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

type Alert struct {
	AlertName string    `json:"alertname"`
	Severity  string    `json:"severity"`
	Instance  string    `json:"instance"`
	Summary   string    `json:"summary"`
	Status    string    `json:"status"`
	StartsAt  time.Time `json:"startsAt"`
}
