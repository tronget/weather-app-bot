package client

import (
	"github.com/tronget/weather-app-bot/internal/config"
	"github.com/tronget/weather-app-bot/internal/weather/service"
)

func GetCorrectCityName(cityName string, cfg *config.Config) (string, error) {
	return service.GetCorrectCityName(cityName, cfg)
}
