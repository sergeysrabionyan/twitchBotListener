package russia

import (
	"twitchBotListener/internal/config"
)

type Russia struct {
	text *string
}

func New(conf *config.Config) *Russia {
	return &Russia{
		text: &conf.Commands.GoolMessage,
	}
}

func (p *Russia) GetCommandAlias() string {
	return "!gool"
}

func (p *Russia) GetValue() string {
	return *p.text
}
