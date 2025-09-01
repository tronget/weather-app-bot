package service

import (
	"fmt"
	"github.com/tronget/weather-app-bot/config"
	"github.com/tronget/weather-app-bot/server"
	"github.com/tronget/weather-app-bot/weather/models"
)

func GetWeatherInfo(cityName string, cfg *config.Config, lang string) (*models.Weather, error) {
	weather, err := server.GetWeatherInfo(cityName, cfg, lang)
	if err != nil {
		return nil, fmt.Errorf("getting weather info during request: %w", err)
	}

	return weather, nil
}
