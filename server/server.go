package server

import (
	"encoding/json"
	"fmt"
	"github.com/tronget/weather-app-bot/config"
	"github.com/tronget/weather-app-bot/pkg/api"
	"github.com/tronget/weather-app-bot/weather/models"
	"io"
	"net/http"
)

func GetCities(cityName string, cfg *config.Config) ([]models.City, error) {
	url := api.RequestCityCoordinatesURL(cityName, cfg)

	jsonBytes, err := getResponseBody(url)
	if err != nil {
		return nil, err
	}

	var data []models.City
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return nil, fmt.Errorf("unmarshaling response body to JSON: %w", err)
	}

	return data, nil
}

func GetWeatherInfo(cityName string, cfg *config.Config, lang string) (*models.Weather, error) {
	url := api.RequestWeatherURL(cityName, cfg, lang)

	jsonBytes, err := getResponseBody(url)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(jsonBytes))

	data := new(models.Weather)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		return nil, fmt.Errorf("unmarshaling weather info to JSON: %w", err)
	}

	return data, nil
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
		return zero, err
	}

	var data T
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		return zero, fmt.Errorf("unmarshaling response body bytes to JSON: %w", err)
	}

	return data, nil
}
