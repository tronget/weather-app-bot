package api

import (
	"fmt"
	config2 "github.com/tronget/weather-app-bot/internal/config"
	"strings"
)

func RequestCityCoordinatesURL(cityName string, cfg *config2.Config) string {
	addressFormat := config2.COORDINATES_REQUEST_URL
	token := cfg.WeatherToken()

	// replace spaces to make correct request
	cityName = strings.ReplaceAll(cityName, " ", "%20")
	address := fmt.Sprintf(addressFormat, cityName, token)
	return address
}

func RequestWeatherURL(cityName string, cfg *config2.Config, lang string) string {
	addressFormat := config2.WEATHER_REQUEST_URL
	token := cfg.WeatherToken()

	// replace spaces to make correct request
	cityName = strings.ReplaceAll(cityName, " ", "%20")
	address := fmt.Sprintf(addressFormat, cityName, token, lang)
	return address
}
