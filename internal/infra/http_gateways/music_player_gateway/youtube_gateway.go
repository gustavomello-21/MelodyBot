package musicplayergateway

import (
	"fmt"

	"github.com/gustavomello-21/melody-bot/internal/entities"
	"github.com/gustavomello-21/melody-bot/internal/usecases/gateways"
)

type YoutubeGateway struct {
	botprefix string
}

func NewYoutubeGateway(botprefix string) gateways.MusicPlayerGateway {
	return &YoutubeGateway{
		botprefix: botprefix,
	}
}

func (y *YoutubeGateway) SearchMusic() (*entities.Video, error) {
	fmt.Println("Searching music...")
	video := entities.Video{
		Title:    "test",
		Url:      "test",
		Duration: 1.2,
	}

	return &video, nil
}

func (y *YoutubeGateway) GetMusics() ([]entities.Video, error) {
	fmt.Println("Getting musics...")
	return nil, nil
}

func (y *YoutubeGateway) PlayMusic() error {
	fmt.Println("Playing music...")
	return nil
}

func (y *YoutubeGateway) StopMusic() error {
	fmt.Println("Stopping music...")
	return nil
}

func (y *YoutubeGateway) AddMusic(video entities.Video) error {
	fmt.Println("Adding music...")
	return nil
}

func (y *YoutubeGateway) RemoveMusic() error {
	fmt.Println("Removing music...")
	return nil
}
