package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetResponseBody(url string) ([]byte, error) {
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

func GetData[T any](url string) (T, error) {
	var zero T

	jsonBytes, err := GetResponseBody(url)
	if err != nil {
		return zero, fmt.Errorf("%s getting response body: %w", url, err)
	}

	var data T
	if err = json.Unmarshal(jsonBytes, &data); err != nil {
		return zero, fmt.Errorf("%s unmarshaling response body bytes to JSON: %w", url, err)
	}

	return data, nil
}
