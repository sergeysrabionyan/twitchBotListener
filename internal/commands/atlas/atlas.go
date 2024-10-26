package atlas

import (
	"errors"
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

func (p *Atlas) Validate() error {
	if p.url == nil || *p.url == "" {
		return errors.New("atlas URL is required")
	}
	return nil
}
