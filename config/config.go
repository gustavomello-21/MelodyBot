package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var (
	Token     string
	BotPrefix string
)

type config struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botprefix"`
}

func LoadConfig() error {
	fmt.Println("Loading config...")

	file, err := os.Open("config.json")
	if err != nil {
		return err
	}

	defer file.Close()

	byte, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	var config config
	if err := json.Unmarshal(byte, &config); err != nil {
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	fmt.Println("Config loaded successfully!")
	return nil
}
