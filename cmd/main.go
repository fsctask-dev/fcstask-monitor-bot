package main

import (
	"context"
	tgbot "fcstask-monitor-bot/internal/bot"
	config "fcstask-monitor-bot/internal/config"
	database "fcstask-monitor-bot/internal/db"
	server "fcstask-monitor-bot/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("[%s][ERROR]: %v", time.Now(), err)
	}

	database.InitDB()

	bot, err := tgbot.NewBot(ctx, cfg)
	if err != nil {
		log.Fatalf("[%s][ERROR]: %v", time.Now(), err)
	}
	bot.Start(ctx)

	serverFiber := server.NewServer(ctx, bot)
	if err := serverFiber.Run(ctx, cfg); err != nil {
		log.Fatalf("[%s][ERROR]: %v", time.Now(), err)
	}
}
