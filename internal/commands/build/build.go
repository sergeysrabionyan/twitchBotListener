package build

import (
	"twitchBotListener/internal/config"
)

type Build struct {
	url *string
}

func New(conf *config.Config) *Build {
	return &Build{
		url: &conf.Commands.BuildURL,
	}
}

func (p *Build) GetCommandAlias() string {
	return "!build"
}

func (p *Build) GetValue() string {
	return *p.url
}
