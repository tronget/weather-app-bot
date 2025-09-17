package config

import (
	"fmt"
	"github.com/tronget/weather-app-bot/internal/network/db"
	"github.com/tronget/weather-app-bot/internal/network/db/models"
	"os"
)

type Config struct {
	telegramToken string
	weatherToken  string
	Users         map[string]*models.User
}

func (cfg *Config) TelegramToken() string {
	return cfg.telegramToken
}

func (cfg *Config) WeatherToken() string {
	return cfg.weatherToken
}

// UserLanguage returns language code from cache.
// In case of cache-miss, it gets language code from database,
// then updates cache and returns received language code.
func (cfg *Config) UserLanguage(username string) string {
	userConfig, ok := cfg.Users[username]
	if ok {
		return userConfig.LangCode
	}

	langCode := db.GetUserLanguage(username)
	cfg.Users[username] = models.NewUser(username, langCode)

	return langCode
}

func (cfg *Config) SetUserLanguage(username string, langCode string) {
	if userConfig, ok := cfg.Users[username]; !ok {
		cfg.Users[username] = models.NewUser(username, langCode)
	} else {
		userConfig.LangCode = langCode
	}
}

func Load() (*Config, error) {
	cfg := &Config{
		Users: make(map[string]*models.User), // Initialize the Users map
	}

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
