package service

import (
	"fmt"
	"github.com/tronget/weather-app-bot/internal/config"
	"github.com/tronget/weather-app-bot/internal/ierrors"
	"github.com/tronget/weather-app-bot/internal/network/api"
)

func GetCorrectCityName(cityName string, cfg *config.Config) (string, error) {
	cities, err := api.GetCities(cityName, cfg)
	if err != nil {
		return "", fmt.Errorf("getting correct city name during request: %w", err)
	}

	if len(cities) == 0 {
		return "", ierrors.NewCityNotFoundError(cityName)
	}

	city := cities[0]

	return city.Name, nil
}
