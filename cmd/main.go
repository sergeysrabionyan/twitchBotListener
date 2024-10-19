package main

import (
	"fmt"
	"log"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/joho/godotenv"

	"twitchBot/internal/commands/atlas"
	"twitchBot/internal/commands/profile"
	"twitchBot/internal/commands/smile"
	"twitchBot/internal/config"
	"twitchBot/internal/services/command_builder"
)

func main() {
	godotenv.Load(".env")

	conf, err := config.NewConfigFromENV()
	if err != nil {
		panic(err)
	}
	commandsService := command_builder.NewService(conf)
	commandsService.RegisterCommands(
		profile.New(),
		atlas.New(),
		smile.New(),
	)

	processMessages(conf, commandsService)
}

func processMessages(conf *config.Config, commandsService *command_builder.Service) {
	client := twitch.NewClient(conf.Twitch.BotUsername, conf.Twitch.BotAccessKey)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Printf("%s: %s\n", message.User.DisplayName, message.Message)

		text, err := commandsService.ProcessCommand(message.Message)
		if err != nil {
			fmt.Println("error while processing message:", err)
			return
		}
		if text == "" {
			return
		}

		client.Say(conf.Twitch.ChannelName, text)
	})

	client.Join(conf.Twitch.ChannelName)

	err := client.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to Twitch IRC: %v", err)
	}

	fmt.Println("Bot is running. Press CTRL+C to exit.")
	select {}
}
