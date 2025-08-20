package config

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tronget/weather-app-bot/internal"
	"os"
)

func GetUpdate() tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	return updateConfig
}

func GetTelegramToken() (string, error) {
	token := os.Getenv(internal.TELEGRAM_APIKEY)
	if token == "" {
		return "", errors.New("TELEGRAM_APIKEY not set")
	}
	return token, nil
}

func GetWeatherToken() (string, error) {
	token := os.Getenv(internal.OPENWEATHER_APIKEY)
	if token == "" {
		return "", errors.New("OPENWEATHER_APIKEY not set")
	}
	return token, nil
}
