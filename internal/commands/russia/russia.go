package russia

import (
	"errors"
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

func (p *Russia) Validate() error {
	if p.text == nil || *p.text == "" {
		return errors.New("russia URL is required")
	}
	return nil
}
