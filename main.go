package main

import (
	"fmt"

	"github.com/gustavomello-21/melody-bot/config"
	messageappgateway "github.com/gustavomello-21/melody-bot/internal/infra/http_gateways/message_app_gateway"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	disc := messageappgateway.NewDiscordGateway(config.Token, config.BotPrefix)
	disc.Start()

	fmt.Print("Bot is running...")
	<-make(chan struct{})
}
