package gateways

type MessageAppGateway interface {
	Start() error
	// SendMessage sends a message
	SendMessage() error

	// SendEmbedMessage sends an embed message
	SendEmbedMessage() error

	// Add Bot in channel
	EnterChannel() error

	// Remove Bot from channel
	LeaveChannel() error

	// GetMessages returns a list of messages
	GetMessages() error
}
