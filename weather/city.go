package weather

import (
	"encoding/json"
	"fmt"
	"github.com/tronget/weather-app-bot/config"
	"github.com/tronget/weather-app-bot/internal"
	"io"
	"net/http"
	"strings"
)

type City struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

func GetCity(cityName string) ([]City, error) {
	url, err := RequestCoordinatesURL(cityName)
	if err != nil {
		return nil, fmt.Errorf("getting coordinates of city: %v", err)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("sending GET request to take coordinates: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status code is not OK: %s", resp.Status)
	}

	jsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}

	var data []City
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling response body: %v", err)
	}

	return data, nil
}

func RequestCoordinatesURL(cityName string) (string, error) {
	addressFormat := internal.COORDINATES_REQUEST_URL
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
