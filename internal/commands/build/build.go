package build

import (
	"errors"
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

func (p *Build) Validate() error {
	if p.url == nil || *p.url == "" {
		return errors.New("build URL is required")
	}
	return nil
}
