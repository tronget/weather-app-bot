package models

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
