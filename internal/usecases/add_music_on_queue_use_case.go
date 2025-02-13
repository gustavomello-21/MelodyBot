package usecases

import (
	"github.com/gustavomello-21/melody-bot/internal/usecases/gateways"
)

type AddMusicOnQueueUseCase struct {
	MusicPlayerGateway gateways.MusicPlayerGateway
}

func NewAddMusicOnQueueUseCase(musicPlayerGateway gateways.MusicPlayerGateway) *AddMusicOnQueueUseCase {
	return &AddMusicOnQueueUseCase{
		MusicPlayerGateway: musicPlayerGateway,
	}
}

func (a *AddMusicOnQueueUseCase) Execute() error {
	video, err := a.MusicPlayerGateway.SearchMusic()
	if err != nil {
		return err
	}

	err = a.MusicPlayerGateway.AddMusic(video)
	if err != nil {
		return err
	}

	return nil
}
