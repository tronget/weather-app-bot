package service

import (
	"fmt"
	"github.com/tronget/weather-app-bot/internal/config"
	"github.com/tronget/weather-app-bot/internal/ierrors"
	"github.com/tronget/weather-app-bot/internal/network/api"
	"github.com/tronget/weather-app-bot/internal/network/server"
	"github.com/tronget/weather-app-bot/internal/weather/models"
)

func GetCorrectCityName(cityName string, cfg *config.Config) (string, error) {
	cities, err := GetCities(cityName, cfg)
	if err != nil {
		return "", fmt.Errorf("getting correct city name during request: %w", err)
	}

	if len(cities) == 0 {
		return "", ierrors.NewCityNotFoundError(cityName)
	}

	city := cities[0]

	return city.Name, nil
}

func GetCities(cityName string, cfg *config.Config) ([]models.City, error) {
	url := api.RequestCityCoordinatesURL(cityName, cfg)

	data, err := server.GetData[[]models.City](url)
	if err != nil {
		return nil, err
	}

	return data, nil
}
