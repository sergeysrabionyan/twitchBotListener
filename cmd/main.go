package main

import (
	"fmt"
	"log"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/joho/godotenv"

	handlers "twitchBotListener/internal/command_handler"
	"twitchBotListener/internal/commands/atlas"
	"twitchBotListener/internal/commands/build"
	"twitchBotListener/internal/commands/profile"
	"twitchBotListener/internal/commands/russia"
	"twitchBotListener/internal/commands/smile"
	"twitchBotListener/internal/config"
	"twitchBotListener/internal/services/command_builder"
)

func main() {
	godotenv.Load(".env")

	// Инициализация конфгиа
	conf, err := config.NewConfigFromENV()
	if err != nil {
		panic(err)
	}
	client := twitch.NewClient(conf.Twitch.BotUsername, conf.Twitch.BotAccessKey)
	commandsService := command_builder.NewService(conf)

	// Регистрация используемых команд
	commandsService.RegisterCommands(
		profile.New(conf),
		atlas.New(conf),
		smile.New(conf),
		build.New(conf),
		russia.New(conf),
	)

	// Установка обработчиков для чата
	client.OnPrivateMessage(handlers.HandleChatMessage(client, conf, commandsService))
	client.OnWhisperMessage(handlers.HandleWhisperMessage(client, conf))

	// Подключение бота в чат
	client.Join(conf.Twitch.ChannelName)

	err = client.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to Twitch IRC: %v", err)
	}

	fmt.Println("Bot is running. Press CTRL+C to exit.")
	select {}
}
