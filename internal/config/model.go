package config

type Config struct {
	BotToken   string `env:"BOT_TOKEN,required"`
	ServerPort string `env:"SERVER_PORT,required"`
	PublicURL  string `env:"PUBLIC_URL,required"`
}
