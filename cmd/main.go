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
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to create config: %v", err)
	}
	
	database.InitDB()
	
	bot, err := tgbot.NewBot(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to create bot: %v", err)
	}
	bot.Start(ctx)
	
	serverFiber := server.NewServer(ctx, bot)
	if err := serverFiber.Run(ctx, cfg); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
