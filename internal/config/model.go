package config

type Config struct {
	Token      string `yaml:"BOT_TOKEN"`
	URL        string `yaml:"URL"`
	ServerPort string `yaml:"SERVER_PORT"`
}
