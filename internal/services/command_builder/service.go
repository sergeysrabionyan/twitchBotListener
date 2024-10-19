package command_builder

import (
	"errors"
	"strings"
	"twitchBot/internal/config"
)

type command interface {
	GetCommandAlias() string
	GetValue() string
	FillData(config *config.Config)
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
		lCmd.FillData(s.CommandData)
		s.commands[cmd.GetCommandAlias()] = lCmd
	}
}

func (s *Service) ProcessCommand(alias string) (string, error) {
	alias = strings.ReplaceAll(removeInvalidUTF8Bytes(alias), " ", "")
	comm, ok := s.commands[alias]
	if !ok {
		return "", errors.New("command not found")
	}
	return comm.GetValue(), nil
}

// Костыль из-за того, что твич иногда присылает не utf-8, а какой-то мусор в конце строки
func removeInvalidUTF8Bytes(input string) string {
	var result []byte
	for _, v := range input {
		if v < 32 || v > 127 {
			continue
		}
		result = append(result, byte(v))
	}
	return string(result)
}
