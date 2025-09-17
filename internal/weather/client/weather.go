package client

import (
	"fmt"
	"github.com/tronget/weather-app-bot/internal/config"
	"github.com/tronget/weather-app-bot/internal/weather/models"
	"github.com/tronget/weather-app-bot/internal/weather/service"
)

func GetWeatherInfo(cityName string, cfg *config.Config, lang string) (*models.Weather, error) {
	weather, err := service.GetWeatherInfo(cityName, cfg, lang)
	if err != nil {
		return nil, fmt.Errorf("getting weather info during request: %w", err)
	}

	return weather, nil
}
