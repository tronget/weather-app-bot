package models

import (
	"encoding/json"
	"fmt"
	"github.com/tronget/weather-app-bot/locales"
	"time"
)

type Weather struct {
	DescriptionList []Description `json:"weather"`
	CityName        string        `json:"name"`
	TimeZone        int           `json:"timezone"`
	Temperature     Temperature   `json:"main"`
	Wind            Wind          `json:"wind"`
	Sys             Sys           `json:"sys"`
}

type Description struct {
	Description string `json:"description"`
	IconID      string `json:"icon"`
}

type Temperature struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
}

type Wind struct {
	Speed float32 `json:"speed"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise time.Time
	Sunset  time.Time
}

type internalSys struct {
	Country     string `json:"country"`
	SunriseUnix int64  `json:"sunrise"`
	SunsetUnix  int64  `json:"sunset"`
}

func (s *Sys) UnmarshalJSON(b []byte) error {

	var tempSys internalSys
	if err := json.Unmarshal(b, &tempSys); err != nil {
		return err
	}

	s.Country = tempSys.Country

	s.Sunrise = time.Unix(tempSys.SunriseUnix, 0).UTC()
	s.Sunset = time.Unix(tempSys.SunsetUnix, 0).UTC()

	return nil
}

func (weather *Weather) BuildMessage(lang string) string {
	weatherDesc := locales.Translate(locales.NO_DATA, lang)
	weatherEmoji := "🌤️"
	if len(weather.DescriptionList) > 0 {
		weatherDesc = weather.DescriptionList[0].Description
		weatherEmoji = iconIDToEmoji(weather.DescriptionList[0].IconID)
	}

	loc := time.FixedZone("local", weather.TimeZone)
	sunrise := weather.Sys.Sunrise.In(loc).Format("15:04")
	sunset := weather.Sys.Sunset.In(loc).Format("15:04")

	msg := fmt.Sprintf(
		locales.WEATHER_MSG_FORMAT,
		weather.CityName, weather.Sys.Country,
		weatherEmoji, weatherDesc,
		locales.Translate(locales.TEMPERATURE, lang),
		weather.Temperature.Temp,
		locales.Translate(locales.FEELS_LIKE, lang),
		weather.Temperature.FeelsLike,
		locales.Translate(locales.WIND, lang),
		weather.Wind.Speed,
		locales.Translate(locales.SUNRISE, lang),
		sunrise,
		locales.Translate(locales.SUNSET, lang),
		sunset,
	)

	return msg
}

func iconIDToEmoji(code string) string {
	switch code {
	case "01d":
		return "☀️"
	case "01n":
		return "🌑"
	case "02d":
		return "🌤️"
	case "02n":
		return "☁️🌙"
	case "03d", "03n":
		return "☁️"
	case "04d", "04n":
		return "☁️☁️"
	case "09d", "09n":
		return "🌧️"
	case "10d":
		return "🌦️"
	case "10n":
		return "🌧️🌙"
	case "11d", "11n":
		return "⛈️"
	case "13d", "13n":
		return "❄️"
	case "50d", "50n":
		return "🌫️"
	default:
		return "❔"
	}
}
