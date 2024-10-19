package build

import (
	"twitchBot/internal/config"
)

type Build struct {
	url *string
}

func New() *Build {
	return &Build{
		url: new(string),
	}
}

func (p *Build) GetCommandAlias() string {
	return "!build"
}

func (p *Build) GetValue() string {
	return *p.url
}

func (p *Build) FillData(conf *config.Config) {
	*p.url = conf.Commands.BuildURL
}
