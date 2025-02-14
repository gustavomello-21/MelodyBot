package gateways

import "github.com/gustavomello-21/melody-bot/internal/entities"

type MessageAppGateway interface {
	Start() error
	// SendMessage sends a message
	SendMessage() error

	// SendEmbedMessage sends an embed message
	SendEmbedMessage(entities.Video) error

	// Add Bot in channel
	EnterChannel(string, string) error

	// Remove Bot from channel
	LeaveChannel() error

	// GetMessages returns a list of messages
	GetMessages() error
}
