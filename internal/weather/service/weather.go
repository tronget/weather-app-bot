package service

import (
	"github.com/tronget/weather-app-bot/internal/config"
	"github.com/tronget/weather-app-bot/internal/network/api"
	"github.com/tronget/weather-app-bot/internal/network/server"
	"github.com/tronget/weather-app-bot/internal/weather/models"
)

func GetWeatherInfo(cityName string, cfg *config.Config, lang string) (*models.Weather, error) {
	url := api.RequestWeatherURL(cityName, cfg, lang)

	data, err := server.GetData[models.Weather](url)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
