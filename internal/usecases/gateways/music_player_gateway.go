package gateways

import "github.com/gustavomello-21/melody-bot/internal/entities"

type MusicPlayerGateway interface {
	// SearchMusic searches for a music
	SearchMusic() (*entities.Video, error)

	// GetMusics returns a list of music
	GetMusics() ([]entities.Video, error)

	// PlayMusic plays a music
	PlayMusic() error

	// StopMusic stops the music
	StopMusic() error

	// AddMusic adds a music to the list
	AddMusic(entities.Video) error

	// RemoveMusic removes a music from the list
	RemoveMusic() error
}
