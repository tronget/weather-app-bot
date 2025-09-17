package api

import (
	"encoding/json"
	"fmt"
	"github.com/tronget/weather-app-bot/internal/config"
	"github.com/tronget/weather-app-bot/internal/weather/models"
	"io"
	"net/http"
)

func GetCities(cityName string, cfg *config.Config) ([]models.City, error) {
	url := RequestCityCoordinatesURL(cityName, cfg)

	data, err := getData[[]models.City](url)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetWeatherInfo(cityName string, cfg *config.Config, lang string) (*models.Weather, error) {
	url := RequestWeatherURL(cityName, cfg, lang)

	data, err := getData[models.Weather](url)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func getResponseBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("sending GET request %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status code is not OK: %s", resp.Status)
	}

	jsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}
	return jsonBytes, nil
}

func getData[T any](url string) (T, error) {
	var zero T

	jsonBytes, err := getResponseBody(url)
	if err != nil {
		return zero, fmt.Errorf("%s getting response body: %w", url, err)
	}

	var data T
	if err = json.Unmarshal(jsonBytes, &data); err != nil {
		return zero, fmt.Errorf("%s unmarshaling response body bytes to JSON: %w", url, err)
	}

	return data, nil
}
