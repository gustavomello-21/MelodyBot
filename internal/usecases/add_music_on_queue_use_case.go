package usecases

import (
	"fmt"

	"github.com/gustavomello-21/melody-bot/internal/usecases/gateways"
)

type AddMusicOnQueueUseCase struct {
	MusicPlayerGateway gateways.MusicPlayerGateway
	MessageAppGateway  gateways.MessageAppGateway
}

func NewAddMusicOnQueueUseCase(musicPlayerGateway gateways.MusicPlayerGateway, messageAppGateway gateways.MessageAppGateway) *AddMusicOnQueueUseCase {
	return &AddMusicOnQueueUseCase{
		MusicPlayerGateway: musicPlayerGateway,
		MessageAppGateway:  messageAppGateway,
	}
}

func (a *AddMusicOnQueueUseCase) Execute(guildId, channelId string) error {
	err := a.MessageAppGateway.EnterChannel(guildId, channelId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	video, err := a.MusicPlayerGateway.SearchMusic()
	if err != nil {
		return err
	}

	err = a.MusicPlayerGateway.AddMusic(*video)
	if err != nil {
		return err
	}

	err = a.MessageAppGateway.SendEmbedMessage(*video)
	if err != nil {
		return err
	}

	return nil
}
