package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Commands struct {
		ProfileURL   string `envconfig:"PROFILE_URL"`
		AtlasURL     string `envconfig:"ATLAS_URL"`
		SmileMessage string `envconfig:"SMILE_MESSAGE"`
		BuildURL     string `envconfig:"BUILD_URL"`
	}
	Twitch struct {
		BotAccessKey string `envconfig:"BOT_ACCESS_KEY"`
		BotUsername  string `envconfig:"BOT_USERNAME"`
		ChannelName  string `envconfig:"CHANNEL_NAME"`
	}
}

func NewConfigFromENV() (*Config, error) {
	var c Config

	envconfig.MustProcess("", &c)

	return &c, nil
}
