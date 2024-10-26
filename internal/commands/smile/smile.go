package smile

import (
	"errors"
	"twitchBotListener/internal/config"
)

type Smile struct {
	text *string
}

func New(conf *config.Config) *Smile {
	return &Smile{
		text: &conf.Commands.SmileMessage,
	}
}

func (p *Smile) GetCommandAlias() string {
	return "!zachem"
}

func (p *Smile) GetValue() string {
	return *p.text
}

func (p *Smile) Validate() error {
	if p.text == nil || *p.text == "" {
		return errors.New("smile URL is required")
	}
	return nil
}
