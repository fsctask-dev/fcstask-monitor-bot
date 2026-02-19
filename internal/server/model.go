package server

import "github.com/gofiber/fiber/v2"

type Server struct {
	app *fiber.App
}

type AlertPayload struct {
	Status string `json:"status"`
	Alerts []struct {
		Annotations map[string]string `json:"annotations"`
		Labels      map[string]string `json:"labels"`
	} `json:"alerts"`
}
