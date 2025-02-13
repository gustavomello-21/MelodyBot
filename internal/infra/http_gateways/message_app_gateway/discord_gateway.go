package messageappgateway

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gustavomello-21/melody-bot/internal/usecases/gateways"
)

type DiscordGateway struct {
	botprefix         string
	discordBotSession *discordgo.Session
}

func NewDiscordGateway(token, botprefix string) gateways.MessageAppGateway {
	discordSession, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil
	}

	return &DiscordGateway{
		botprefix:         botprefix,
		discordBotSession: discordSession,
	}
}

func (d *DiscordGateway) Start() error {
	err := d.discordBotSession.Open()
	if err != nil {
		return err
	}

	return nil
}

func (d *DiscordGateway) SendMessage() error {
	message, err := d.discordBotSession.ChannelMessageSend("1339093294621266033", "hello world")
	if err != nil {
		return err
	}

	fmt.Println("Message sent successfully: ", message.ID)

	return nil
}

func (d *DiscordGateway) SendEmbedMessage() error {
	return nil
}

func (d *DiscordGateway) EnterChannel() error {
	return nil
}

func (d *DiscordGateway) LeaveChannel() error {
	return nil
}

func (d *DiscordGateway) GetMessages() error {
	return nil
}
