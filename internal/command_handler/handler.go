package command_handler

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v4"

	"twitchBotListener/internal/config"
	"twitchBotListener/internal/services/command_builder"
)

func HandleChatMessage(client *twitch.Client, conf *config.Config, commandsService *command_builder.Service) func(message twitch.PrivateMessage) {
	return func(message twitch.PrivateMessage) {
		fmt.Printf("%s: %s\n", message.User.DisplayName, message.Message)

		text, err := commandsService.ProcessCommand(message.Message)
		if err != nil {
			fmt.Println("error while processing message:", err)
			return
		}
		if text == "" {
			return
		}

		client.Reply(conf.Twitch.ChannelName, message.ID, text)
	}
}
