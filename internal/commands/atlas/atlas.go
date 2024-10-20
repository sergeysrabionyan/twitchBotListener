package atlas

import (
	"twitchBotListener/internal/config"
)

type Atlas struct {
	url *string
}

func New(config *config.Config) *Atlas {
	return &Atlas{
		url: &config.Commands.AtlasURL,
	}
}

func (p *Atlas) GetCommandAlias() string {
	return "!atlas"
}

func (p *Atlas) GetValue() string {
	return *p.url
}
