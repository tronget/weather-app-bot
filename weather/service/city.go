package service

import (
	"fmt"
	"github.com/tronget/weather-app-bot/config"
	"github.com/tronget/weather-app-bot/ierrors"
	"github.com/tronget/weather-app-bot/server"
)

func GetCorrectCityName(cityName string, cfg *config.Config) (string, error) {
	cities, err := server.GetCities(cityName, cfg)
	if err != nil {
		return "", fmt.Errorf("getting correct city name during request: %w", err)
	}

	if len(cities) == 0 {
		return "", ierrors.NewCityNotFoundError(cityName)
	}

	city := cities[0]

	return city.Name, nil
}
