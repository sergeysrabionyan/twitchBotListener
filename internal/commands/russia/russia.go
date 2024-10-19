package russia

import (
	"twitchBotListener/internal/config"
)

type Russia struct {
	text *string
}

func New() *Russia {
	return &Russia{
		text: new(string),
	}
}

func (p *Russia) GetCommandAlias() string {
	return "!gool"
}

func (p *Russia) GetValue() string {
	return *p.text
}

func (p *Russia) FillData(conf *config.Config) {
	*p.text = conf.Commands.GoolMessage
}
