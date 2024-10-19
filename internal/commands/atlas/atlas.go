package atlas

import (
	"twitchBotListener/internal/config"
)

type Atlas struct {
	url *string
}

func New() *Atlas {
	return &Atlas{
		url: new(string),
	}
}

func (p *Atlas) GetCommandAlias() string {
	return "!atlas"
}

func (p *Atlas) GetValue() string {
	return *p.url
}

func (p *Atlas) FillData(config *config.Config) {
	*p.url = config.Commands.AtlasURL
}
