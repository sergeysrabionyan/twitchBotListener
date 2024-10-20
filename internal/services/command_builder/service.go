package command_builder

import (
	"errors"
	"strings"
	"twitchBotListener/internal/config"
	"twitchBotListener/internal/utls"
)

type command interface {
	GetCommandAlias() string
	GetValue() string
	Validate() error
}

type Service struct {
	CommandData *config.Config
	commands    map[string]command
}

func NewService(cnf *config.Config) *Service {
	return &Service{
		CommandData: cnf,
		commands:    map[string]command{},
	}
}

func (s *Service) RegisterCommands(com ...command) {
	for _, cmd := range com {
		lCmd := cmd
		s.commands[cmd.GetCommandAlias()] = lCmd
	}
}

func (s *Service) ProcessCommand(alias string) (string, error) {
	comm, ok := s.commands[s.formatString(alias)]
	if !ok {
		return "", errors.New("command not found")
	}
	return comm.GetValue(), nil
}

func (s *Service) formatString(alias string) string {
	return strings.ToLower(strings.ReplaceAll(utls.RemoveInvalidUTF8Bytes(alias), " ", ""))
}
