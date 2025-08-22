package config

import (
	"fmt"
	"os"
)

type Config struct {
	telegramToken string
	weatherToken  string
}

func (c Config) TelegramToken() string {
	return c.telegramToken
}

func (c Config) WeatherToken() string {
	return c.weatherToken
}

func Load() (*Config, error) {
	cfg := &Config{}

	var err error
	cfg.telegramToken, err = getToken(TELEGRAM_APIKEY)
	if err != nil {
		return nil, fmt.Errorf("getting Telegram API token: %v", err)
	}

	cfg.weatherToken, err = getToken(OPENWEATHER_APIKEY)
	if err != nil {
		return nil, fmt.Errorf("getting OpenWeather API token: %v", err)
	}

	return cfg, nil
}

func getToken(env string) (string, error) {
	token, ok := os.LookupEnv(env)
	if !ok {
		return "", fmt.Errorf("%s not set", env)
	}
	return token, nil
}
