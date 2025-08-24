package config

import (
	"fmt"
	"os"
)

type Config struct {
	telegramToken string
	weatherToken  string
	userLanguage  map[int64]string
}

func (c *Config) TelegramToken() string {
	return c.telegramToken
}

func (c *Config) WeatherToken() string {
	return c.weatherToken
}

func (c *Config) UserLanguage(id int64) string {
	return c.userLanguage[id]
}

func (c *Config) SetUserLanguage(id int64, lang string) {
	c.userLanguage[id] = lang
}

func Load() (*Config, error) {
	cfg := &Config{userLanguage: make(map[int64]string)}

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
