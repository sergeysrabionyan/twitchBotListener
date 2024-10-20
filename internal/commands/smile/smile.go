package smile

import (
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
