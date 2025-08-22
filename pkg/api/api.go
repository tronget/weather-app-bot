package api

import (
	"fmt"
	"github.com/tronget/weather-app-bot/config"
	"strings"
)

func RequestCityCoordinatesURL(cityName string, cfg *config.Config) string {
	addressFormat := config.COORDINATES_REQUEST_URL
	token := cfg.WeatherToken()

	// replace spaces to make correct request
	cityName = strings.ReplaceAll(cityName, " ", "%20")
	address := fmt.Sprintf(addressFormat, cityName, token)
	return address
}

func RequestWeatherURL(cityName string, cfg *config.Config) string {
	addressFormat := config.WEATHER_REQUEST_URL
	token := cfg.WeatherToken()

	// replace spaces to make correct request
	cityName = strings.ReplaceAll(cityName, " ", "%20")
	address := fmt.Sprintf(addressFormat, cityName, token)
	return address
}
