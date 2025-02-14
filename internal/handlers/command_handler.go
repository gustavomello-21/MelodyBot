package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gustavomello-21/melody-bot/internal/usecases"
	"github.com/gustavomello-21/melody-bot/internal/usecases/gateways"
)

type CommandHandler struct {
	commands           map[string]func(s *discordgo.Session, m *discordgo.MessageCreate)
	MessageAppGateway  gateways.MessageAppGateway
	MusicPlayerGateway gateways.MusicPlayerGateway
}

func NewCommandHandler(messageAppGateway gateways.MessageAppGateway, musicPlayerGateway gateways.MusicPlayerGateway) *CommandHandler {
	handler := &CommandHandler{
		commands:           make(map[string]func(s *discordgo.Session, m *discordgo.MessageCreate)),
		MessageAppGateway:  messageAppGateway,
		MusicPlayerGateway: musicPlayerGateway,
	}

	handler.Register("play", handler.playCommand)

	return handler
}

func (ch *CommandHandler) Register(
	command string,
	handlerFunc func(s *discordgo.Session, m *discordgo.MessageCreate),
) {
	ch.commands[command] = handlerFunc
}

func (ch *CommandHandler) Handler(s *discordgo.Session, m *discordgo.MessageCreate, botprefix string) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == botprefix+"play" {
		ch.playCommand(s, m)
	}
}

func (ch *CommandHandler) playCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	addMusicOnQueuUseCase := usecases.NewAddMusicOnQueueUseCase(ch.MusicPlayerGateway, ch.MessageAppGateway)

	if err := addMusicOnQueuUseCase.Execute(); err != nil {
		return
	}
}
