package profile

import (
	"twitchBotListener/internal/config"
)

type Profile struct {
	url *string
}

func New() *Profile {
	return &Profile{
		url: new(string),
	}
}

func (p *Profile) GetCommandAlias() string {
	return "!profile"
}

func (p *Profile) GetValue() string {
	return *p.url
}

func (p *Profile) FillData(config *config.Config) {
	*p.url = config.Commands.ProfileURL
}
