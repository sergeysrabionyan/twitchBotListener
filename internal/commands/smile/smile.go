package smile

import (
	"twitchBot/internal/config"
)

type Smile struct {
	image *string
}

func New() *Smile {
	return &Smile{
		image: new(string),
	}
}

func (p *Smile) GetCommandAlias() string {
	return "!zachem"
}

func (p *Smile) GetValue() string {
	return *p.image
}

func (p *Smile) FillData(conf *config.Config) {
	*p.image = conf.Commands.SmileMessage
}
