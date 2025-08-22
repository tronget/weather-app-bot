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

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("sending GET request to take city info: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status code is not OK: %s", resp.Status)
	}

	jsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}

	var data []models.City
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling response body to JSON: %v", err)
	}

	return data, nil
}
