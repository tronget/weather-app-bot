package config

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

func GetUpdate() tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	return updateConfig
}

func GetToken() string {
	return os.Getenv("TELEGRAM_TOKENAPI")
}
