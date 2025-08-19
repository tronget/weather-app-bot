package config

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

func GetUpdate() tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	return updateConfig
}

func GetToken() (string, error) {
	token := os.Getenv("TELEGRAM_TOKENAPI")
	if token == "" {
		return "", errors.New("TELEGRAM_TOKENAPI not set")
	}
	return token, nil
}
