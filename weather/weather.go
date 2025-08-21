package weather

import (
	"fmt"
	"github.com/tronget/weather-app-bot/config"
	"github.com/tronget/weather-app-bot/internal"
	"strings"
)

type Main struct {
	WeatherList []Weather `json:"weather"`
	Name        string    `json:"name"`
	Temperature `json:"main"`
}

type Weather struct {
	Description string `json:"description"`
	IconID      string `json:"icon"`
}

type Temperature struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:""`
}

func GetWeather(city string) error {
	return nil
}

func RequestWeatherURL(cityName string) (string, error) {
	addressFormat := internal.WEATHER_REQUEST_URL
	token, err := config.GetWeatherToken()
	if err != nil {
		return "", err
	}

	// replace spaces to make correct request
	cityName = strings.ReplaceAll(cityName, " ", "%20")
	fmt.Println(cityName)
	address := fmt.Sprintf(addressFormat, cityName, token)
	return address, nil
}

func IconIdToPngPath(iconID string) string {
	return ""
}
