package messageappgateway

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gustavomello-21/melody-bot/internal/entities"
	"github.com/gustavomello-21/melody-bot/internal/handlers"
)

type DiscordGateway struct {
	botprefix         string
	discordBotSession *discordgo.Session
	CommandHandler    *handlers.CommandHandler
}

func NewDiscordGateway(token, botprefix string) *DiscordGateway {
	discordSession, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil
	}

	discordGateway := &DiscordGateway{
		botprefix:         botprefix,
		discordBotSession: discordSession,
	}

	return discordGateway
}

func (d *DiscordGateway) Start() error {
	d.discordBotSession.AddHandler(d.handleMessage)
	d.discordBotSession.Identify.Intents = discordgo.IntentGuildMessages
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

func (d *DiscordGateway) SendEmbedMessage(video entities.Video) error {
	a := discordgo.MessageEmbed{
		Title:       "Now Playing...",
		Description: fmt.Sprintf("A música que está tocando é %s", video.Title),
		Color:       100,
	}
	d.discordBotSession.ChannelMessageSendEmbed("1339093294621266033", &a)
	fmt.Println("Sending embed message...")
	return nil
}

func (d *DiscordGateway) EnterChannel(guildId, channelId string) error {
	_, err := d.discordBotSession.ChannelVoiceJoin(guildId, "879851686393413668", true, false)
	if err != nil {
		return err
	}
	fmt.Println("Entering channel....")
	return nil
}

func (d *DiscordGateway) LeaveChannel() error {
	fmt.Println("Leavening channel....")
	return nil
}

func (d *DiscordGateway) GetMessages() error {
	fmt.Println("Geting messages")
	return nil
}

func (d *DiscordGateway) handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	d.CommandHandler.Handler(s, m, d.botprefix)
}
