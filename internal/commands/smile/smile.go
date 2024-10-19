package smile

import (
	"twitchBot/internal/config"
)

type Smile struct {
	text *string
}

func New() *Smile {
	return &Smile{
		text: new(string),
	}
}

func (p *Smile) GetCommandAlias() string {
	return "!zachem"
}

func (p *Smile) GetValue() string {
	return *p.text
}

func (p *Smile) FillData(conf *config.Config) {
	*p.text = conf.Commands.SmileMessage
}
