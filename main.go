package main

import (
	"fmt"

	"github.com/gustavomello-21/melody-bot/config"
	"github.com/gustavomello-21/melody-bot/internal/handlers"
	messageappgateway "github.com/gustavomello-21/melody-bot/internal/infra/http_gateways/message_app_gateway"
	musicplayergateway "github.com/gustavomello-21/melody-bot/internal/infra/http_gateways/music_player_gateway"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Bot is running...")

	disc := messageappgateway.NewDiscordGateway(config.Token, config.BotPrefix)
	yt := musicplayergateway.NewYoutubeGateway(config.BotPrefix)

	commandHandler := handlers.NewCommandHandler(disc, yt)
	disc.CommandHandler = commandHandler

	disc.Start()

	<-make(chan struct{})
}
