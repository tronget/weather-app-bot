package config

import (
	"errors"
	"github.com/tronget/weather-app-bot/internal"
	"os"
)

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
