package service

import (
	"fmt"
	"github.com/tronget/weather-app-bot/internal/config"
	"github.com/tronget/weather-app-bot/internal/network/api"
	"github.com/tronget/weather-app-bot/internal/weather/models"
)

func GetWeatherInfo(cityName string, cfg *config.Config, lang string) (*models.Weather, error) {
	weather, err := api.GetWeatherInfo(cityName, cfg, lang)
	if err != nil {
		return nil, fmt.Errorf("getting weather info during request: %w", err)
	}

	return weather, nil
}
