package profile

import (
	"errors"
	"twitchBotListener/internal/config"
)

type Profile struct {
	url *string
}

func New(config *config.Config) *Profile {
	return &Profile{
		url: &config.Commands.ProfileURL,
	}
}

func (p *Profile) GetCommandAlias() string {
	return "!profile"
}

func (p *Profile) GetValue() string {
	return *p.url
}

func (p *Profile) Validate() error {
	if p.url == nil || *p.url == "" {
		return errors.New("profile URL is required")
	}
	return nil
}
